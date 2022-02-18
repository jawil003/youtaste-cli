package services

import (
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/services/db"
	"encoding/json"
	"os"
)

type ScrapperService struct {
}

func (_ ScrapperService) ScrapUrlAndOpeningTimes(scrapper models.Scrapper, highestPoll models.PollWithCount) {
	page, err := scrapper.OpenInNewBrowserAndJoin()
	if err != nil {
		return
	}

	page, err = scrapper.Login(os.Getenv(db.YoutastePhone), os.Getenv(db.YoutastePassword), page)
	if err != nil {
		return
	}

	page, err = scrapper.SearchForRestaurant(highestPoll.RestaurantName, page)
	if err != nil {
		return
	}

	url, err := scrapper.GetUrl(page)
	if err != nil {
		return
	}

	err = DB().Settings().Create(db.RestaurantUrl, *url)
	if err != nil {
		return
	}

	err = os.Setenv(db.RestaurantUrl, *url)
	if err != nil {
		return
	}

	openingTimes, err := scrapper.GetOpeningTimes(page)
	if err != nil {
		return
	}

	openingString, err := json.Marshal(openingTimes)

	err = DB().Settings().Create(db.OpeningTimes, string(openingString))
	if err != nil {
		return
	}

	err = os.Setenv(db.OpeningTimes, string(openingString))
	if err != nil {
		return
	}
}
