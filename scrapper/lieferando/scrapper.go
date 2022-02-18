package lieferando

import (
	"bs-to-scrapper/server/datastructures"
	"bs-to-scrapper/server/models"
	"errors"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type LieferandoScrapper struct {
	models.Scrapper
}

func (_ LieferandoScrapper) Login(_, _ string, page *rod.Page) (*rod.Page, error) {
	return page, nil
}

func (_ LieferandoScrapper) OpenInNewBrowserAndJoin() (*rod.Page, error) {
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

func (_ LieferandoScrapper) SearchForRestaurant(name string, page *rod.Page) (*rod.Page, error) {
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

func (_ LieferandoScrapper) GetOpeningTimes(page *rod.Page) (*datastructures.Weekdays, error) {

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

	if len(elements) < 13 {
		return nil, errors.New("no opening times found")
	}

	openingTimes := datastructures.Weekdays{}

	openingTimes.Monday, err = elements[1].Text()

	if err != nil {
		return nil, err
	}

	openingTimes.Tuesday, err = elements[3].Text()

	if err != nil {
		return nil, err
	}

	openingTimes.Wednesday, err = elements[5].Text()

	if err != nil {
		return nil, err
	}

	openingTimes.Thursday, err = elements[7].Text()

	if err != nil {
		return nil, err
	}

	openingTimes.Friday, err = elements[9].Text()

	if err != nil {
		return nil, err
	}

	openingTimes.Saturday, err = elements[11].Text()

	if err != nil {
		return nil, err
	}

	openingTimes.Sunday, err = elements[13].Text()

	if err != nil {
		return nil, err
	}

	return &openingTimes, nil

}
