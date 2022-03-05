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
	page, err := browser.Page(proto.TargetCreateTarget{URL: "https://www.lieferando.de/lieferservice/essen/arnsberg-dortmund-44269"})
	if err != nil {
		errorLogger.Printf("Error while opening page: %s", err)
		return nil, err
	}

	return page, nil
}

func (_ Scrapper) SearchForRestaurant(name string, page *rod.Page) (*rod.Page, error) {

	err := page.Wait(nil, `
			() => {
				return document.querySelectorAll('div[data-qa=\"restaurant-card\"]')?.length > 0
			}
`, nil)
	if err != nil {
		errorLogger.Printf("Error while waiting for restaurant card: %s", err)
		return nil, err
	}

	element, err := page.Element("input[type=search]")
	if err != nil {
		errorLogger.Printf("Error while getting search input: %s", err)
		return nil, err
	}

	err = element.Input(name)
	if err != nil {
		errorLogger.Printf("Error while inputting search: %s", err)
		return nil, err
	}

	wait := page.WaitRequestIdle(time.Duration(10)*time.Second, []string{"/search"}, []string{})
	wait()

	element, err = page.Element("a[data-qa=link]")

	if err != nil {
		return nil, err
	}

	eval, err := page.Eval("window.location.href")
	if err != nil {
		errorLogger.Printf("Error while getting url: %s", err)
		return nil, err
	}

	oldUrl := eval.Value.Str()

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

	return page, nil

}

func (_ Scrapper) GetOpeningTimes(page *rod.Page) (*datastructures.Weekdays, error) {

	button, err := page.Element("*[role=\"button\"][data-qa=\"restaurant-header-action-info\"]")
	if err != nil {
		errorLogger.Printf("Error while getting opening times button: %s", err)
		return nil, err
	}

	err = button.Click(proto.InputMouseButtonLeft)
	if err != nil {
		errorLogger.Printf("Error while clicking opening times button: %s", err)
		return nil, err
	}

	err = page.Wait(nil, `()=>{
			const element = document.querySelector("div[data-qa=\"restaurant-info-modal-overlay\"]");
			return !!element;
		}`, nil)
	if err != nil {
		errorLogger.Printf("Error while waiting for opening times modal: %s", err)
		return nil, err
	}

	err = page.Wait(nil, `()=>{
			const elements = document.querySelectorAll("*[data-qa=\"restaurant-info-modal-info-shipping-times-element-element\"] *[data-qa=\"text\"]");
			return elements && elements.length > 0;
		}`, nil)
	if err != nil {
		errorLogger.Printf("Error while waiting for opening times: %s", err)
		return nil, err
	}

	elements, err := page.Elements("*[data-qa=\"restaurant-info-modal-info-shipping-times-element-element\"] *[data-qa=\"text\"]")
	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	if len(elements) < 13 {
		err = errors.New("no opening times found")
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	openingTimes := datastructures.Weekdays{}

	openingTimes.Monday, err = elements[1].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	openingTimes.Tuesday, err = elements[3].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	openingTimes.Wednesday, err = elements[5].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	openingTimes.Thursday, err = elements[7].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	openingTimes.Friday, err = elements[9].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	openingTimes.Saturday, err = elements[11].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	openingTimes.Sunday, err = elements[13].Text()

	if err != nil {
		errorLogger.Printf("Error while getting opening times: %s", err)
		return nil, err
	}

	return &openingTimes, nil

}

func (_ Scrapper) GetUrl(page *rod.Page) (*string, error) {
	rem, err := page.Eval("window.location.href")
	if err != nil {
		errorLogger.Printf("Error while getting url: %s", err)
		return nil, err
	}

	res := rem.Value.Str()

	return &res, nil
}

func (_ Scrapper) GoToUrl(url string, page *rod.Page) (*rod.Page, error) {
	err := page.Navigate(url)
	if err != nil {
		errorLogger.Printf("Error while navigating to url %s: %s", url, err)
		return nil, err
	}

	return page, nil
}

func (_ Scrapper) SelectProduct(name string, variants []string, page *rod.Page) (*rod.Page, error) {

	searchToggleButton, err := page.Element("*[data-qa=\"menu-category-nav-categories-action-search\"]")
	if err != nil {
		errorLogger.Printf("Error while getting search toggle button: %s", err)
		return nil, err
	}

	err = searchToggleButton.Click(proto.InputMouseButtonLeft)
	if err != nil {
		errorLogger.Printf("Error while clicking search toggle button: %s", err)
		return nil, err
	}

	inputSelector := "input[type=\"search\"]"

	err = page.Wait(nil, fmt.Sprintf(`() => {
				const element = document.querySelector("%s");
				return Boolean(element);
			}`, inputSelector), nil)
	if err != nil {
		errorLogger.Printf("Error while waiting for search input: %s", err)
		return nil, err
	}

	inputSearch, err := page.Element(inputSelector)

	if err != nil {
		errorLogger.Printf("Error while getting search input: %s", err)
		return nil, err
	}

	err = inputSearch.Input(name)
	if err != nil {
		errorLogger.Printf("Error while inputing search input: %s", err)
		return nil, err
	}

	item, err := page.ElementR("*", name)
	if err != nil {
		errorLogger.Printf("Error while getting item: %s", err)
		return nil, err
	}

	err = item.Click(proto.InputMouseButtonLeft)
	if err != nil {
		errorLogger.Printf("Error while clicking item: %s", err)
		return nil, err
	}

	return page, nil
}
