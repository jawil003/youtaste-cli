package db

import (
	"bs-to-scrapper/server/enums"
	"bs-to-scrapper/server/models"
	"encoding/json"
	"errors"
	"github.com/thoas/go-funk"
	bolt "go.etcd.io/bbolt"
	"strconv"
)

type PollService struct {
}

func (_ ServiceCollection) Poll() PollService {
	return PollService{}
}

func (_ PollService) Create(poll models.Poll, user string, provider string) error {
	db, err := OpenDbConnection()
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucketPollsCount := tx.Bucket([]byte(enums.PollsCount))
		bucketPollsByUser := tx.Bucket([]byte(enums.PollsUser))
		bucketPollsByProvider := tx.Bucket([]byte(enums.PollsProvider))
		bucketPollsByUrl := tx.Bucket([]byte(enums.PollsUrl))

		pollsByUserString := bucketPollsByUser.Get([]byte(user))

		var pollsByUserStringUnmarshal []string

		if pollsByUserString != nil {

			err := json.Unmarshal(pollsByUserString, &pollsByUserStringUnmarshal)

			if err != nil {

				return err
			}
		}

		hasAlreadyVoted := funk.ContainsString(pollsByUserStringUnmarshal, poll.RestaurantName)

		if hasAlreadyVoted {
			return errors.New("you have already voted for this restaurant")
		} else {
			pollsString := bucketPollsCount.Get([]byte(poll.RestaurantName))

			if pollsString != nil {
				pollsCount, err := strconv.Atoi(string(pollsString))

				if err != nil {
					return err
				}

				err = bucketPollsCount.Put([]byte(poll.RestaurantName), []byte(strconv.Itoa(pollsCount+1)))

				if err != nil {

					return err
				}

				return nil

			} else {
				err = bucketPollsCount.Put([]byte(poll.RestaurantName), []byte("1"))

				if provider == "" {
					return errors.New("provider is empty")
				}

				err = bucketPollsByProvider.Put([]byte(poll.RestaurantName), []byte(provider))

				if err != nil {
					return err
				}

				if err != nil {

					return err
				}
			}

			pollsByUserStringUnmarshal = append(pollsByUserStringUnmarshal, poll.RestaurantName)

			pollsByUserString, err = json.Marshal(pollsByUserStringUnmarshal)

			if err != nil {
				return err
			}

			err := bucketPollsByUser.Put([]byte(user), pollsByUserString)
			if err != nil {
				return err
			}

			err = bucketPollsByUrl.Put([]byte(poll.RestaurantName), []byte(poll.Url))
			if err != nil {
				return err
			}
		}

		return nil

	})
	if err != nil {
		return err
	}
	return nil
}

func (_ PollService) GetAll() ([]models.PollWithCount, error) {
	db, err := OpenDbConnection()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	pollsWithCount := make([]models.PollWithCount, 0)

	err = db.View(func(tx *bolt.Tx) error {
		bucketPollsCount := tx.Bucket([]byte(enums.PollsCount))
		bucketPollsByProvider := tx.Bucket([]byte(enums.PollsProvider))
		bucketPollsByUrl := tx.Bucket([]byte(enums.PollsUrl))

		cursor := bucketPollsCount.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {

			provider := bucketPollsByProvider.Get(k)

			url := bucketPollsByUrl.Get(k)

			poll := models.Poll{
				RestaurantName: string(k),
				Provider:       string(provider),
				Url:            string(url),
			}

			stringCOnvInt, err := strconv.Atoi(string(v))

			if err != nil {
				return err
			}

			pollsWithCount = append(pollsWithCount, models.PollWithCount{
				Poll:  poll,
				Count: stringCOnvInt,
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return pollsWithCount, nil
}

func (ps PollService) PersistFinalResult() (*models.PollWithCount, error) {
	polls, err := ps.GetAll()

	if err != nil {
		return nil, err
	}

	var highestPoll *models.PollWithCount

	for _, poll := range polls {
		if highestPoll == nil {
			highestPoll = &poll
		} else {
			if poll.Count > highestPoll.Count {
				highestPoll = &poll
			}
		}
	}

	str, err := json.Marshal(highestPoll)

	if err != nil {
		return nil, err
	}

	err = SettingsService{}.Create(enums.ChoosenRestaurant, string(str))
	if err != nil {
		return nil, err
	}

	return highestPoll, nil

}
