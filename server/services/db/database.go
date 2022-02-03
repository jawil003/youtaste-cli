package db

import (
	bolt "go.etcd.io/bbolt"
)

type ServiceCollection struct{}

func OpenDbConnection() (*bolt.DB, error) {

	open, err := bolt.Open("tastyfood.db", 0666, nil)
	if err != nil {
		return nil, err
	}
	return open, nil
}
