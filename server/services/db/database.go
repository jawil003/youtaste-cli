package db

import (
	bolt "go.etcd.io/bbolt"
)

const (
	Orders        = "orders"
	PollsCount    = "polls_count"
	PollsUser     = "polls_user"
	PollsProvider = "polls_provider"
	Settings      = "settings"
	PollsUrl      = "polls_url"
)

type ServiceCollection struct{}

func OpenDbConnection() (*bolt.DB, error) {

	open, err := bolt.Open("tastyfood.db", 0666, nil)
	if err != nil {
		return nil, err
	}

	err = open.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(Orders))
		_, err = tx.CreateBucketIfNotExists([]byte(PollsCount))
		_, err = tx.CreateBucketIfNotExists([]byte(PollsUser))
		_, err = tx.CreateBucketIfNotExists([]byte(PollsProvider))
		_, err = tx.CreateBucketIfNotExists([]byte(Settings))
		_, err = tx.CreateBucketIfNotExists([]byte(PollsUrl))

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
