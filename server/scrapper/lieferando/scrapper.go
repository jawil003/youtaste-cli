package lieferando

import (
	"bs-to-scrapper/server/datastructures"
	"bs-to-scrapper/server/models"
	"errors"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"time"
)

type Scrapper struct {
	models.Scrapper
}

func (_ Scrapper) Login(_, _ string, page *rod.Page) (*rod.Page, error) {
	return page, nil
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
	page, err := browser.Page(proto.TargetCreateTarget{URL: "https://www.lieferando.de/lieferservice/essen/arnsberg-dortmund-44269"})
	if err != nil {
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
		return nil, err
	}

	element, err := page.Element("input[type=search]")
	if err != nil {
		return nil, err
	}

	err = element.Input(name)
	if err != nil {
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
		return nil, err
	}

	oldUrl := eval.Value.Str()

	err = element.Click(proto.InputMouseButtonLeft)
	if err != nil {
		return nil, err
	}

	err = page.Wait(nil, `(oldUrl) => {
		return window.location.href !== oldUrl;
	}`, []interface{}{oldUrl})

	if err != nil {
		return nil, err
	}

	return page, nil

}

func (_ Scrapper) GetOpeningTimes(page *rod.Page) (*datastructures.Weekdays, error) {

	//FIXME: Fix getting Opening Times for Lieferando

	button, err := page.Element("*[role=\"button\"][data-qa=\"restaurant-header-action-info\"]")
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

func (_ Scrapper) GetUrl(page *rod.Page) (*string, error) {
	rem, err := page.Eval("window.location.href")
	if err != nil {
		return nil, err
	}

	res := rem.Value.Str()

	return &res, nil
}
