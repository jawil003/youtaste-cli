package services

import (
	"strconv"
	"time"
)

type TimeService struct {
}

func (_ TimeService) ConvertOpeningTimeToDate(openingTime string) (*time.Time, error) {

	hours := openingTime[0:2]
	minutes := openingTime[3:5]

	hoursNumber, err := strconv.Atoi(hours)

	if err != nil {
		return nil, err
	}

	minutesNumber, err := strconv.Atoi(minutes)

	if err != nil {
		return nil, err
	}

	t := time.Now()

	t = time.Date(t.Year(), t.Month(), t.Day(), hoursNumber, minutesNumber, 0, 0, t.Location())

	return &t, nil

}
