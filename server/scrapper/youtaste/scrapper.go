package youtaste

import (
	"bs-to-scrapper/server/datastructures"
	"bs-to-scrapper/server/models"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type Scrapper struct {
	models.Scrapper
}

func (_ Scrapper) OpenInCurrentBrowserAndJoin() *rod.Page {
	u := launcher.NewUserMode().
		MustLaunch()

	controller := rod.New().ControlURL(u)

	page := controller.MustConnect().NoDefaultDevice().MustPage("https://youtaste.com/")

	return page
}

func (_ Scrapper) OpenInNewBrowserAndJoin(headless bool) (*rod.Page, error) {
	var browser *rod.Browser

	if !headless {
		u, err := launcher.New().Headless(false).Launch()

		if err != nil {
			return nil, err
		}

		browser = rod.New().ControlURL(u)
	} else {
		browser = rod.New()
	}

	err := browser.Connect()
	if err != nil {
		return nil, err
	}
	page, err := browser.Page(proto.TargetCreateTarget{URL: "https://youtaste.com/"})
	if err != nil {
		return nil, err
	}

	return page, nil
}

func (_ Scrapper) Login(phoneNumber, password string, page *rod.Page) (*rod.Page, error) {

	element, err := page.ElementR("#navigation a", "Einloggen")
	if err != nil {
		return nil, err
	}

	wait := page.WaitNavigation(proto.PageLifecycleEventNameNetworkAlmostIdle)
	err = element.Click(proto.InputMouseButtonLeft)
	if err != nil {
		return nil, err
	}

	wait()

	phoneInput, err := page.Element("input[placeholder=\"Telefonnummer\"]")
	if err != nil {
		return nil, err
	}

	err = phoneInput.Input(phoneNumber)
	if err != nil {
		return nil, err
	}

	passwordInput, err := page.Element("input[placeholder=\"Passwort\"]")
	if err != nil {
		return nil, err
	}

	err = passwordInput.Input(password)
	if err != nil {
		return nil, err
	}

	loginButton, err := page.ElementR("button", "Anmelden")
	if err != nil {
		return nil, err
	}

	wait = page.WaitNavigation(proto.PageLifecycleEventNameNetworkAlmostIdle)

	err = loginButton.Click(proto.InputMouseButtonLeft)
	if err != nil {
		return nil, err
	}

	wait()

	return page, nil

}

func (_ Scrapper) SearchForRestaurant(name string, page *rod.Page) (*rod.Page, error) {

	searchInput, err := page.Element("input#search-restaurant-input")
	if err != nil {
		return nil, err
	}
	err = searchInput.Input(name)
	if err != nil {
		return nil, err
	}
	err = searchInput.Press(input.Enter)
	if err != nil {
		return nil, err
	}

	resterauntElement, err := page.Element("#restaurantList a")
	if err != nil {
		return nil, err
	}

	err = resterauntElement.Click(proto.InputMouseButtonLeft)

	if err != nil {
		return nil, err
	}

	err = page.WaitLoad()
	if err != nil {
		return nil, err
	}

	return page, nil
}

func (_ Scrapper) SelectProduct(name string, variants []string, page *rod.Page) error {
	element, err := page.ElementR("#search-content-div a", name)
	if err != nil {
		return err
	}

	err = element.Click(proto.InputMouseButtonLeft)

	for _, variant := range variants {
		regex := fmt.Sprintf("/\\s*%s\\s*/gmi", variant)
		element, err = page.ElementR("#productModalForm div.text-black", regex)
		if err != nil {
			return err
		}
		err := element.Click(proto.InputMouseButtonLeft)
		if err != nil {
			return err
		}
	}

	submitBtn, err := page.Element("input[type=\"submit\"]")
	if err != nil {
		return err
	}

	err = submitBtn.Click(proto.InputMouseButtonLeft)
	if err != nil {
		return err
	}

	return nil
}

func (_ Scrapper) GetOpeningTimes(page *rod.Page) (*datastructures.Weekdays, error) {

	weekdays := datastructures.Weekdays{}

	openingTimesElements, err := page.Elements("div#openhours li")

	if err != nil {
		return nil, err
	}

	weekdays.Monday = openingTimesElements[0].MustText()
	weekdays.Tuesday = openingTimesElements[1].MustText()
	weekdays.Wednesday = openingTimesElements[2].MustText()
	weekdays.Thursday = openingTimesElements[3].MustText()
	weekdays.Friday = openingTimesElements[4].MustText()
	weekdays.Saturday = openingTimesElements[5].MustText()
	weekdays.Sunday = openingTimesElements[6].MustText()

	return &weekdays, nil
}

func (_ Scrapper) GetUrl(page *rod.Page) (*string, error) {
	rem, err := page.Eval("window.location.href")
	if err != nil {
		return nil, err
	}

	res := rem.Value.Str()

	return &res, nil
}
