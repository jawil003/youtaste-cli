package lieferando

import (
	"bs-to-scrapper/server/datastructures"
	"bs-to-scrapper/server/logger"
	"bs-to-scrapper/server/models"
	"errors"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"time"
)

var infoLogger = logger.Logger().Info
var errorLogger = logger.Logger().Error

type Scrapper struct {
	models.ScrapUrlAndOpeningTimesScrapper
}

func (_ Scrapper) Login(_, _ string, page *rod.Page) (*rod.Page, error) {
	infoLogger.Println("User is logged in")
	return page, nil
}

func (_ Scrapper) OpenInNewBrowserAndJoin(headless bool) (*rod.Page, error) {

	var browser *rod.Browser

	if !headless {
		u, err := launcher.New().Headless(false).Launch()

		if err != nil {
			errorLogger.Printf("Error while launching browser: %s", err)
			return nil, err
		}

		browser = rod.New().ControlURL(u)

		infoLogger.Println("Opening browser in headed mode")

	} else {
		browser = rod.New()

		infoLogger.Println("Opening browser in headless mode")
	}

	browser = browser.Timeout(time.Minute * 5)

	err := browser.Connect()
	if err != nil {
		errorLogger.Printf("Error while connecting to browser: %s", err)
		return nil, err
	}
	infoLogger.Println("Connected to browser")
	page, err := browser.Page(proto.TargetCreateTarget{URL: "https://www.lieferando.de/lieferservice/essen/arnsberg-dortmund-44269"})
	if err != nil {
		errorLogger.Printf("Error while opening page: %s", err)
		return nil, err
	}
	infoLogger.Printf("Opened page %s", "https://www.lieferando.de/lieferservice/essen/arnsberg-dortmund-44269")

	return page, nil
}

func (_ Scrapper) SearchForRestaurant(name string, page *rod.Page) (*rod.Page, error) {

	selector := "div[data-qa=\"restaurant-card\"]"

	err := page.Wait(nil, `
			(selector) => {
				return document.querySelectorAll(selector)?.length > 0
			}
`, []interface{}{selector})
	if err != nil {
		errorLogger.Printf("Error while waiting for restaurant card: %s", err)
		return nil, err
	}

	infoLogger.Printf("Found restaurant card with selector %s", selector)

	selector = "input[type=search]"

	element, err := page.Element(selector)
	if err != nil {
		errorLogger.Printf("Error while getting search input: %s", err)
		return nil, err
	}

	infoLogger.Printf("Found search input with selector %s", selector)

	err = element.Input(name)
	if err != nil {
		errorLogger.Printf("Error while inputting search: %s", err)
		return nil, err
	}

	infoLogger.Printf("Inputted search %s", name)

	wait := page.WaitRequestIdle(time.Duration(10)*time.Second, []string{"/search"}, []string{})
	wait()

	infoLogger.Println("Waited for /search to be idle")

	selector = "a[data-qa=link]"

	element, err = page.Element(selector)

	if err != nil {
		return nil, err
	}

	infoLogger.Println("Found link %s", selector)

	eval, err := page.Eval("window.location.href")
	if err != nil {
		errorLogger.Printf("Error while getting url: %s", err)
		return nil, err
	}

	oldUrl := eval.Value.Str()

	infoLogger.Printf("Got current url %s", oldUrl)

	err = element.Click(proto.InputMouseButtonLeft)
	if err != nil {
		errorLogger.Printf("Error while clicking link: %s", err)
		return nil, err
	}

	err = page.Wait(nil, `(oldUrl) => {
		return window.location.href !== oldUrl;
	}`, []interface{}{oldUrl})

	if err != nil {
		errorLogger.Printf("Error while waiting for url change: %s", err)
		return nil, err
	}

	infoLogger.Println("Wait for url change success")

	return page, nil

}

func (_ Scrapper) GetOpeningTimes(page *rod.Page) (*datastructures.Weekdays, error) {

	selector := "*[role=\"button\"][data-qa=\"restaurant-header-action-info\"]"

	button, err := page.Element(selector)
	if err != nil {
		errorLogger.Printf("Error while getting opening times button: %s", err)
		return nil, err
	}

	infoLogger.Printf("Found opening times button with selector %s", selector)

	err = button.Click(proto.InputMouseButtonLeft)
	if err != nil {
		errorLogger.Printf("Error while clicking opening times button: %s", err)
		return nil, err
	}

	infoLogger.Println("Clicked opening times button")

	selector = "div[data-qa=\"restaurant-info-modal-overlay\"]"

	err = page.Wait(nil, `(selector)=>{
			const element = document.querySelector(selector);
			return !!element;
		}`, []interface{}{selector})
	if err != nil {
		errorLogger.Printf("Error while waiting for opening times modal: %s", err)
		return nil, err
	}

	infoLogger.Printf("Waited for opening times modal with selector %s", selector)

	selector = "*[data-qa=\"restaurant-info-modal-info-shipping-times-element-element\"] *[data-qa=\"text\"]"

	err = page.Wait(nil, `(selector)=>{
			const elements = document.querySelectorAll(selector);
			return elements && elements.length > 0;
		}`, []interface{}{selector})
	if err != nil {
		errorLogger.Printf("Error while waiting for opening times elements: %s", err)
		return nil, err
	}

	infoLogger.Printf("Waited for opening times elements with selector %s", selector)

	elements, err := page.Elements("*[data-qa=\"restaurant-info-modal-info-shipping-times-element-element\"] *[data-qa=\"text\"]")
	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got opening times elements with selector %s", selector)

	if len(elements) < 13 {
		err = errors.New("no opening times found")
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	infoLogger.Printf("Found opening times elements with selector %s", selector)

	openingTimes := datastructures.Weekdays{}

	openingTimes.Monday, err = elements[1].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got opening times for Monday: %s", openingTimes.Monday)

	openingTimes.Tuesday, err = elements[3].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got opening times for Tuesday: %s", openingTimes.Tuesday)

	openingTimes.Wednesday, err = elements[5].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got opening times for Wednesday: %s", openingTimes.Wednesday)

	openingTimes.Thursday, err = elements[7].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got opening times for Thursday: %s", openingTimes.Thursday)

	openingTimes.Friday, err = elements[9].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got opening times for Friday: %s", openingTimes.Friday)

	openingTimes.Saturday, err = elements[11].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got opening times for Saturday: %s", openingTimes.Saturday)

	openingTimes.Sunday, err = elements[13].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got opening times for Sunday: %s", openingTimes.Sunday)

	return &openingTimes, nil

}

func (_ Scrapper) GetUrl(page *rod.Page) (*string, error) {
	rem, err := page.Eval("window.location.href")
	if err != nil {
		errorLogger.Printf("Error while getting url: %s", err)
		return nil, err
	}

	res := rem.Value.Str()

	infoLogger.Printf("Got url: %s", res)

	return &res, nil
}

func (_ Scrapper) GoToUrl(url string, page *rod.Page) (*rod.Page, error) {
	err := page.Navigate(url)
	if err != nil {
		errorLogger.Printf("Error while navigating to url %s: %s", url, err)
		return nil, err
	}

	infoLogger.Printf("Navigated to url %s", url)

	return page, nil
}

func (_ Scrapper) SelectProduct(name string, variants []string, page *rod.Page) (*rod.Page, error) {

	selector := "*[data-qa=\"menu-category-nav-categories-action-search\"]"

	searchToggleButton, err := page.Element(selector)
	if err != nil {
		errorLogger.Printf("Error while getting search toggle button: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got search toggle button with selector %s", selector)

	err = searchToggleButton.Click(proto.InputMouseButtonLeft)
	if err != nil {
		errorLogger.Printf("Error while clicking search toggle button: %s", err)
		return nil, err
	}

	infoLogger.Printf("Clicked search toggle button with selector %s", selector)

	inputSelector := "input[type=\"search\"]"

	err = page.Wait(nil, fmt.Sprintf(`() => {
				const element = document.querySelector("%s");
				return Boolean(element);
			}`, inputSelector), nil)
	if err != nil {
		errorLogger.Printf("Error while waiting for search input: %s", err)
		return nil, err
	}

	infoLogger.Printf("Waited for search input with selector %s", inputSelector)

	inputSearch, err := page.Element(inputSelector)

	if err != nil {
		errorLogger.Printf("Error while getting search input: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got search input with selector %s", inputSelector)

	err = inputSearch.Input(name)
	if err != nil {
		errorLogger.Printf("Error while inputing search input: %s", err)
		return nil, err
	}

	infoLogger.Printf("Inputed search input with selector %s", inputSelector)

	item, err := page.ElementR("*", name)
	if err != nil {
		errorLogger.Printf("Error while getting item: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got item with selector %s", name)

	err = item.Click(proto.InputMouseButtonLeft)
	if err != nil {
		errorLogger.Printf("Error while clicking item: %s", err)
		return nil, err
	}

	infoLogger.Printf("Clicked item with selector %s", name)

	return page, nil
}
