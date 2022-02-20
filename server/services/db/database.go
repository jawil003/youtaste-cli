package db

import (
	"bs-to-scrapper/server/enums"
	bolt "go.etcd.io/bbolt"
)

type ServiceCollection struct{}

func OpenDbConnection() (*bolt.DB, error) {

	open, err := bolt.Open("tastyfood.db", 0666, nil)
	if err != nil {
		return nil, err
	}

	err = open.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(enums.Orders))
		_, err = tx.CreateBucketIfNotExists([]byte(enums.PollsCount))
		_, err = tx.CreateBucketIfNotExists([]byte(enums.PollsUser))
		_, err = tx.CreateBucketIfNotExists([]byte(enums.PollsProvider))
		_, err = tx.CreateBucketIfNotExists([]byte(enums.Settings))
		_, err = tx.CreateBucketIfNotExists([]byte(enums.PollsUrl))

		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return open, nil
}
