package lieferando

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func OpenInNewBrowserAndJoinLieferando() (*rod.Page, error) {
	browser := rod.New()

	err := browser.Connect()
	if err != nil {
		return nil, err
	}
	page, err := browser.Page(proto.TargetCreateTarget{URL: "https://www.lieferando.de/lieferservice/essen/arnsberg-dortmund-44269"})
	if err != nil {
		return nil, err
	}

	return page, nil
}

func SearchForRestaurant(name string, page *rod.Page) (*rod.Page, error) {
	element, err := page.Element("input[type=search]")
	if err != nil {
		return nil, err
	}

	err = element.Input(name)
	if err != nil {
		return nil, err
	}

	element, err = page.Element("a[data-qa=link]")

	if err != nil {
		return nil, err
	}

	err = element.Click(proto.InputMouseButtonLeft)
	if err != nil {
		return nil, err
	}

	err = page.WaitLoad()
	if err != nil {
		return nil, err
	}

	return page, nil

}

func GetOpeningTimes(page *rod.Page) ([]string, error) {

	button, err := page.Element("button[\"role=button\"][data-qa=\"restaurant-header-action-info\"]")
	if err != nil {
		return nil, err
	}

	err = button.Click(proto.InputMouseButtonLeft)
	if err != nil {
		return nil, err
	}

	err = page.Wait(nil, `()=>{
			const element = document.querySelector("div[data-qa=\"restaurant-info-modal-overlay\"]");
			return !!element;
		}`, nil)
	if err != nil {
		return nil, err
	}

	elements, err := page.Elements("*[data-qa=restaurant-info-modal-info-shipping-times-element-element] *[data-qa=text]")
	if err != nil {
		return nil, err
	}

	var openingTimes []string

	for _, element := range elements {
		text, err := element.Text()
		if err != nil {
			return nil, err
		}
		openingTimes = append(openingTimes, text)
	}

	return openingTimes, nil

}
