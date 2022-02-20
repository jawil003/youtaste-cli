package services

import (
	"bs-to-scrapper/server/enums"
	"bs-to-scrapper/server/models"
	"encoding/json"
	"os"
)

type ScrapperService struct {
}

func (_ ScrapperService) ScrapUrlAndOpeningTimes(scrapper models.Scrapper, highestPoll models.PollWithCount) {
	page, err := scrapper.OpenInNewBrowserAndJoin(true)
	if err != nil {
		return
	}

	page, err = scrapper.Login(os.Getenv(enums.YoutastePhone), os.Getenv(enums.YoutastePassword), page)
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

	err = DB().Settings().Create(enums.RestaurantUrl, *url)
	if err != nil {
		return
	}

	err = os.Setenv(enums.RestaurantUrl, *url)
	if err != nil {
		return
	}

	openingTimes, err := scrapper.GetOpeningTimes(page)
	if err != nil {
		return
	}

	openingString, err := json.Marshal(openingTimes)

	err = DB().Settings().Create(enums.OpeningTimes, string(openingString))
	if err != nil {
		return
	}

	err = os.Setenv(enums.OpeningTimes, string(openingString))
	if err != nil {
		return
	}
}
