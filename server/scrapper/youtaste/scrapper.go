package youtaste

import (
	"bs-to-scrapper/server/datastructures"
	"bs-to-scrapper/server/logger"
	"bs-to-scrapper/server/models"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"time"
)

var errorLogger = logger.Logger().Error
var infoLogger = logger.Logger().Info

type Scrapper struct {
	models.ScrapUrlAndOpeningTimesScrapper
}

func (_ Scrapper) OpenInCurrentBrowserAndJoin() *rod.Page {
	u := launcher.NewUserMode().
		MustLaunch()

	browser := rod.New().ControlURL(u)

	infoLogger.Println("Opening page in current browser")

	browser = browser.Timeout(time.Minute * 5)

	browser = browser.NoDefaultDevice()

	infoLogger.Println("Set timeoput for page to load %s", time.Minute*5)

	err := browser.Connect()

	if err != nil {
		errorLogger.Println("Error connecting to page %s", err)
	}

	page, err := browser.Page(proto.TargetCreateTarget{URL: "https://youtaste.com/"})

	if err != nil {
		errorLogger.Println("Error opening page %s", err)
	}

	infoLogger.Println("Page loaded")

	return page
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
		infoLogger.Println("Launched browser")
	} else {
		browser = rod.New()
		infoLogger.Println("Launching headless browser")
	}

	browser = browser.Timeout(time.Minute * 5)

	infoLogger.Println("Set timeout for page to load %s", time.Minute*5)

	err := browser.Connect()
	if err != nil {
		errorLogger.Printf("Error while connecting to browser: %s", err)
		return nil, err
	}
	infoLogger.Println("Connected to browser")
	page, err := browser.Page(proto.TargetCreateTarget{URL: "https://youtaste.com/"})
	if err != nil {
		errorLogger.Printf("Error while opening page: %s", err)
		return nil, err
	}

	infoLogger.Println("Page loaded")

	return page, nil
}

func (_ Scrapper) Login(phoneNumber, password string, page *rod.Page) (*rod.Page, error) {

	element, err := page.ElementR("#navigation a", "Einloggen")
	if err != nil {
		errorLogger.Printf("Error while finding element: %s", err)
		return nil, err
	}

	infoLogger.Printf("Found element with selector %s and regex %s", "#navigation a", "Einloggen")

	wait := page.WaitNavigation(proto.PageLifecycleEventNameNetworkAlmostIdle)
	err = element.Click(proto.InputMouseButtonLeft)
	if err != nil {
		errorLogger.Printf("Error while clicking element: %s", err)
		return nil, err
	}

	infoLogger.Println("Clicked element")

	wait()

	infoLogger.Println("Waited for navigation")

	phoneInput, err := page.Element("input[placeholder=\"Telefonnummer\"]")
	if err != nil {
		errorLogger.Printf("Error while finding element: %s", err)
		return nil, err
	}

	infoLogger.Printf("Found element with selector %s", "input[placeholder=\"Telefonnummer\"]")

	err = phoneInput.Input(phoneNumber)
	if err != nil {
		errorLogger.Printf("Error while inputing phone number: %s", err)
		return nil, err
	}

	infoLogger.Printf("Inputted phone number %s", phoneNumber)

	passwordInput, err := page.Element("input[placeholder=\"Passwort\"]")
	if err != nil {
		errorLogger.Printf("Error while finding element: %s", err)
		return nil, err
	}

	infoLogger.Printf("Found element with selector %s", "input[placeholder=\"Passwort\"]")

	err = passwordInput.Input(password)
	if err != nil {
		errorLogger.Printf("Error while inputing password: %s", err)
		return nil, err
	}

	infoLogger.Printf("Inputted password %s", password)

	loginButton, err := page.ElementR("button", "Anmelden")
	if err != nil {
		errorLogger.Printf("Error while finding element: %s", err)
		return nil, err
	}

	infoLogger.Printf("Found element with selector %s and regex %s", "button", "Anmelden")

	wait = page.WaitNavigation(proto.PageLifecycleEventNameNetworkAlmostIdle)

	err = loginButton.Click(proto.InputMouseButtonLeft)
	if err != nil {
		errorLogger.Printf("Error while clicking element: %s", err)
		return nil, err
	}

	infoLogger.Println("Clicked element")

	wait()

	infoLogger.Println("Waited for navigation")

	return page, nil

}

func (_ Scrapper) SearchForRestaurant(name string, page *rod.Page) (*rod.Page, error) {

	searchInput, err := page.Element("input#search-restaurant-input")
	if err != nil {
		errorLogger.Printf("Error while finding element: %s", err)
		return nil, err
	}

	infoLogger.Printf("Found element with selector %s", "input#search-restaurant-input")

	err = searchInput.Input(name)
	if err != nil {
		errorLogger.Printf("Error while inputing name: %s", err)
		return nil, err
	}

	infoLogger.Printf("Inputted name %s", name)

	err = searchInput.Press(input.Enter)
	if err != nil {
		errorLogger.Printf("Error while pressing enter: %s", err)
		return nil, err
	}

	infoLogger.Println("Pressed enter")

	resterauntElement, err := page.Element("#restaurantList a")
	if err != nil {
		errorLogger.Printf("Error while finding element: %s", err)
		return nil, err
	}

	infoLogger.Printf("Found element with selector %s", "#restaurantList a")

	err = resterauntElement.Click(proto.InputMouseButtonLeft)

	if err != nil {
		errorLogger.Printf("Error while clicking element: %s", err)
		return nil, err
	}

	infoLogger.Println("Clicked element")

	err = page.WaitLoad()
	if err != nil {
		errorLogger.Printf("Error while waiting for load: %s", err)
		return nil, err
	}

	infoLogger.Println("Waited for load")

	return page, nil
}

func (_ Scrapper) SelectProduct(name string, variants []string, page *rod.Page) (*rod.Page, error) {
	element, err := page.ElementR("#search-content-div a", name)
	if err != nil {
		errorLogger.Printf("Error while finding element: %s", err)
		return nil, err
	}

	infoLogger.Printf("Found element with selector %s and regex %s", "#search-content-div a", name)

	err = element.Click(proto.InputMouseButtonLeft)

	if err != nil {
		errorLogger.Printf("Error while clicking element: %s", err)
		return nil, err
	}

	infoLogger.Println("Clicked element")

	for _, variant := range variants {
		regex := fmt.Sprintf("/\\s*%s\\s*/gmi", variant)
		element, err = page.ElementR("#productModalForm div.text-black", regex)
		if err != nil {
			errorLogger.Printf("Error while finding element: %s", err)
			return nil, err
		}
		infoLogger.Printf("Found element with selector %s and regex %s", "#productModalForm div.text-black", regex)
		err := element.Click(proto.InputMouseButtonLeft)
		if err != nil {
			errorLogger.Printf("Error while clicking element: %s", err)
			return nil, err
		}

		infoLogger.Println("Clicked element")
	}

	submitBtn, err := page.Element("input[type=\"submit\"]")
	if err != nil {
		errorLogger.Printf("Error while finding element: %s", err)
		return nil, err
	}

	infoLogger.Printf("Found element with selector %s", "input[type=\"submit\"]")

	err = submitBtn.Click(proto.InputMouseButtonLeft)
	if err != nil {
		errorLogger.Printf("Error while clicking element: %s", err)
		return nil, err
	}

	infoLogger.Println("Clicked element")

	return page, nil
}

func (_ Scrapper) GetOpeningTimes(page *rod.Page) (*datastructures.Weekdays, error) {

	weekdays := datastructures.Weekdays{}

	openingTimesElements, err := page.Elements("div#openhours li")

	infoLogger.Printf("Found %d elements with selector %s", len(openingTimesElements), "div#openhours li")

	if err != nil {
		errorLogger.Printf("Error while finding element: %s", err)
		return nil, err
	}

	weekdays.Monday, err = openingTimesElements[0].Text()

	if err != nil {
		errorLogger.Printf("Error while getting text for monday: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got text for monday %s", weekdays.Monday)

	weekdays.Tuesday, err = openingTimesElements[1].Text()

	if err != nil {
		errorLogger.Printf("Error while getting text for tuesday: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got text for tuesday %s", weekdays.Tuesday)

	weekdays.Wednesday, err = openingTimesElements[2].Text()

	if err != nil {
		errorLogger.Printf("Error while getting text for wednesday: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got text for wednesday %s", weekdays.Wednesday)

	weekdays.Thursday, err = openingTimesElements[3].Text()

	if err != nil {
		errorLogger.Printf("Error while getting text for thursday: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got text for thursday %s", weekdays.Thursday)

	weekdays.Friday, err = openingTimesElements[4].Text()

	if err != nil {
		errorLogger.Printf("Error while getting text for friday: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got text for friday %s", weekdays.Friday)

	weekdays.Saturday, err = openingTimesElements[5].Text()

	if err != nil {
		errorLogger.Printf("Error while getting text for saturday: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got text for saturday %s", weekdays.Saturday)

	weekdays.Sunday, err = openingTimesElements[6].Text()

	if err != nil {
		errorLogger.Printf("Error while getting text for sunday: %s", err)
		return nil, err
	}

	infoLogger.Printf("Got text for sunday %s", weekdays.Sunday)

	return &weekdays, nil
}

func (_ Scrapper) GetUrl(page *rod.Page) (*string, error) {
	rem, err := page.Eval("window.location.href")
	if err != nil {
		errorLogger.Printf("Error while getting url: %s", err)
		return nil, err
	}

	res := rem.Value.Str()

	infoLogger.Printf("Got url %s", res)

	return &res, nil
}

func (_ Scrapper) GoToUrl(url string, page *rod.Page) (*rod.Page, error) {
	err := page.Navigate(url)
	if err != nil {
		errorLogger.Printf("Error while navigating to url: %s", err)
		return nil, err
	}

	infoLogger.Printf("Navigated to url %s", url)

	return page, nil
}
