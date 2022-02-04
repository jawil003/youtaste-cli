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

	err = open.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("orders"))

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

func CloseConnection(db *bolt.DB) {
	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	return
}
