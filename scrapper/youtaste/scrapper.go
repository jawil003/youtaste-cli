package youtaste

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

func Login(phoneNumber, password string) (*rod.Page, error) {
	u := launcher.NewUserMode().
		MustLaunch()

	controller := rod.New().ControlURL(u)

	page := controller.MustConnect().NoDefaultDevice().MustPage("https://youtaste.com/")

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

func SearchForRestaurant(name string, page *rod.Page) error {

	searchInput := page.MustElement("input#search-restaurant-input")
	searchInput.MustInput(name)
	searchInput.MustPress(input.Enter)

	err := page.MustElement("#restaurantList a").Click(proto.InputMouseButtonLeft)
	if err != nil {
		return err
	}

	return nil
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
