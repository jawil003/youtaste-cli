package db

import (
	"bs-to-scrapper/server/router/api/ws"
	"encoding/json"
	"github.com/thoas/go-funk"
	bolt "go.etcd.io/bbolt"
)

type PollService struct {
}

func (_ PollService) Create(poll ws.Poll, user string) error {
	db, err := OpenDbConnection()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("polls"))

		pollsString := bucket.Get([]byte(user))

		if len(pollsString) > 0 {
			var polls []ws.Poll

			err := json.Unmarshal(pollsString, &polls)

			if err != nil {
				return err
			}

			exists := funk.Contains(polls, poll)

			if exists {
				return nil
			}

			res, err := json.Marshal(append(polls, poll))

			if err != nil {
				return err
			}

			err = bucket.Put([]byte(user), res)
		} else {
			res, err := json.Marshal([]ws.Poll{poll})

			if err != nil {
				return err
			}

			err = bucket.Put([]byte(user), res)

			if err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}

		return nil

	})
	if err != nil {
		return err
	}

	return nil
}
