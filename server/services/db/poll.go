package db

import (
	"bs-to-scrapper/server/models"
	"errors"
	bolt "go.etcd.io/bbolt"
	"strconv"
)

type PollService struct {
}

func (_ ServiceCollection) Poll() PollService {
	return PollService{}
}

func (_ PollService) Create(poll models.Poll) error {
	db, err := OpenDbConnection()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("polls"))

		pollsString := bucket.Get([]byte(poll.RestaurantName))

		if pollsString != nil {
			return errors.New("poll already exists")
		}

		err := bucket.Put([]byte(poll.RestaurantName), []byte("0"))

		if err != nil {
			return err
		}

		return nil

	})
	if err != nil {
		defer db.Close()
		return err
	}

	defer db.Close()
	return nil
}

func (_ PollService) Choose(poll models.Poll) error {
	db, err := OpenDbConnection()

	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("polls"))

		voteNumber := bucket.Get([]byte(poll.RestaurantName))

		if voteNumber == nil {

			return errors.New("poll does not exist")
		}

		voteNumberInt, err := strconv.Atoi(string(voteNumber))

		if err != nil {
			return err
		}

		err = bucket.Put([]byte(poll.RestaurantName), []byte(strconv.Itoa(voteNumberInt+1)))

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		defer db.Close()
		return err
	}

	defer db.Close()
	return nil
}
