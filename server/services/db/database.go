package db

import (
	bolt "go.etcd.io/bbolt"
)

const (
	ORDERS         = "orders"
	POLLS_COUNT    = "polls_count"
	POLLS_USER     = "polls_user"
	POLLS_PROVIDER = "polls_provider"
	SETTINGS       = "settings"
)

type ServiceCollection struct{}

func OpenDbConnection() (*bolt.DB, error) {

	open, err := bolt.Open("tastyfood.db", 0666, nil)
	if err != nil {
		return nil, err
	}

	err = open.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(ORDERS))
		_, err = tx.CreateBucketIfNotExists([]byte(POLLS_COUNT))
		_, err = tx.CreateBucketIfNotExists([]byte(POLLS_USER))
		_, err = tx.CreateBucketIfNotExists([]byte(POLLS_PROVIDER))
		_, err = tx.CreateBucketIfNotExists([]byte(SETTINGS))

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
