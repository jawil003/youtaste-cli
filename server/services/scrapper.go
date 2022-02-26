package services

import (
	"bs-to-scrapper/server/datastructures"
	"bs-to-scrapper/server/enums"
	"bs-to-scrapper/server/models"
	"encoding/json"
	"os"
)

type ScrapperService struct {
}

func (_ ScrapperService) ScrapUrlAndOpeningTimes(scrapper models.ScrapUrlAndOpeningTimesScrapper, highestPoll models.PollWithCount) (*datastructures.Weekdays, error) {
	page, err := scrapper.OpenInNewBrowserAndJoin(true)
	if err != nil {
		return nil, err
	}

	page, err = scrapper.Login(os.Getenv(enums.YoutastePhone), os.Getenv(enums.YoutastePassword), page)
	if err != nil {
		return nil, err
	}

	page, err = scrapper.SearchForRestaurant(highestPoll.RestaurantName, page)
	if err != nil {
		return nil, err
	}

	url, err := scrapper.GetUrl(page)
	if err != nil {
		return nil, err
	}

	err = DB().Settings().Create(enums.RestaurantUrl, *url)
	if err != nil {
		return nil, err
	}

	err = os.Setenv(enums.RestaurantUrl, *url)
	if err != nil {
		return nil, err
	}

	openingTimes, err := scrapper.GetOpeningTimes(page)
	if err != nil {
		return nil, err
	}

	openingString, err := json.Marshal(openingTimes)

	err = DB().Settings().Create(enums.OpeningTimes, string(openingString))
	if err != nil {
		return nil, err
	}

	err = os.Setenv(enums.OpeningTimes, string(openingString))
	if err != nil {
		return nil, err
	}

	return openingTimes, nil
}

func (_ ScrapperService) OrderMeals(scrapper models.SelectProductScrapper, highestPoll models.PollWithCount, orders []models.Order) error {
	page, err := scrapper.OpenInNewBrowserAndJoin(true)
	if err != nil {
		return err
	}

	page, err = scrapper.Login(os.Getenv(enums.YoutastePhone), os.Getenv(enums.YoutastePassword), page)
	if err != nil {
		return err
	}

	page, err = scrapper.GoToUrl(highestPoll.Url, page)
	if err != nil {
		return err
	}

	for _, order := range orders {
		page, err = scrapper.SelectProduct(order.Name, order.Variants, page)
		if err != nil {
			return err
		}
	}

	return nil

}
