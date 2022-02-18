package models

import (
	"bs-to-scrapper/server/datastructures"
	"github.com/go-rod/rod"
)

type Scrapper interface {
	OpenInNewBrowserAndJoin() (*rod.Page, error)
	Login(phoneNumber, password string, page *rod.Page) (*rod.Page, error)
	SearchForRestaurant(name string, page *rod.Page) (*rod.Page, error)
	GetOpeningTimes(page *rod.Page) (*datastructures.Weekdays, error)
	GetUrl(page *rod.Page) (*string, error)
}
