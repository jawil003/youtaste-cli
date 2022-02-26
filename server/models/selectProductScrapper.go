package models

import (
	"github.com/go-rod/rod"
)

type SelectProductScrapper interface {
	OpenInNewBrowserAndJoin(headless bool) (*rod.Page, error)
	Login(phoneNumber, password string, page *rod.Page) (*rod.Page, error)
	SearchForRestaurant(name string, page *rod.Page) (*rod.Page, error)
	SelectProduct(name string, variants []string, page *rod.Page) (*rod.Page, error)
	GoToUrl(url string, page *rod.Page) (*rod.Page, error)
}
