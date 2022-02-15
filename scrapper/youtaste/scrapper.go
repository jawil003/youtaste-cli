package youtaste

import (
	"bs-to-scrapper/server/datastructures"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func OpenInCurrentBrowserAndJoinYouTaste() *rod.Page {
	u := launcher.NewUserMode().
		MustLaunch()

	controller := rod.New().ControlURL(u)

	page := controller.MustConnect().NoDefaultDevice().MustPage("https://youtaste.com/")

	return page
}

func OpenInNewBrowserAndJoinYouTaste() *rod.Page {
	page := rod.New().MustConnect().MustPage("https://youtaste.com/")

	return page
}

func Login(phoneNumber, password string, page *rod.Page) (*rod.Page, error) {

	err := page.MustElementR("#navigation a", "Einloggen").Click(proto.InputMouseButtonLeft)
	if err != nil {
		return nil, err
	}

	page.MustWaitNavigation()

	phoneInput := page.MustElement("input[placeholder=\"Telefonnummer\"]")

	phoneInput.MustInput(phoneNumber)

	passwordInput := page.MustElement("input[placeholder=\"Passwort\"]")
	passwordInput.MustInput(password)

	err = page.MustElementR("button", "Anmelden").Click(proto.InputMouseButtonLeft)
	if err != nil {
		return nil, err
	}

	page.MustWaitNavigation()

	return page, nil

}

func SearchForRestaurant(name string, page *rod.Page) (*rod.Page, error) {

	searchInput := page.MustElement("input#search-restaurant-input")
	searchInput.MustInput(name)
	searchInput.MustPress(input.Enter)

	resterauntElement := page.MustElement("#restaurantList a")

	err := resterauntElement.Click(proto.InputMouseButtonLeft)

	if err != nil {
		return nil, err
	}

	page.MustWaitLoad()

	return page, nil
}

func SelectProduct(name string, variants []string, page *rod.Page) error {
	element, err := page.ElementR("#search-content-div a", name)
	if err != nil {
		return err
	}

	err = element.Click(proto.InputMouseButtonLeft)

	for _, variant := range variants {
		regex := fmt.Sprintf("/\\s*%s\\s*/gmi", variant)
		element = page.MustElementR("#productModalForm div.text-black", regex)
		err := element.Click(proto.InputMouseButtonLeft)
		if err != nil {
			return err
		}
	}

	page.MustElement("input[type=\"submit\"]").MustClick()

	if err != nil {
		return err
	}

	return nil
}

func GetOpeningTimes(page *rod.Page) (datastructures.Weekdays, error) {

	weekdays := datastructures.Weekdays{}

	page.MustScreenshot("screenshot.png")

	openingTimesElements := page.MustElements("div#openhours li")

	weekdays.Monday = openingTimesElements[0].MustText()
	weekdays.Tuesday = openingTimesElements[1].MustText()
	weekdays.Wednesday = openingTimesElements[2].MustText()
	weekdays.Thursday = openingTimesElements[3].MustText()
	weekdays.Friday = openingTimesElements[4].MustText()
	weekdays.Saturday = openingTimesElements[5].MustText()
	weekdays.Sunday = openingTimesElements[6].MustText()

	return weekdays, nil
}

func GetUrl(page *rod.Page) string {
	return page.MustEval("window.location.href").Str()
}
