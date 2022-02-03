package database

import (
	"errors"
	bolt "go.etcd.io/bbolt"
)

var instance *Database

func Init() *Database {
	if instance == nil {
		instance = &Database{}
	}

	return instance
}

type Database struct {
	db *bolt.DB
}

func (_ Database) Order() Order {
	return Order{}
}

func (d *Database) Init() (*bolt.DB, error) {

	if d.db != nil {
		return nil, errors.New("database is initialized")
	}

	open, err := bolt.Open("tastyfood.db", 0666, nil)
	if err != nil {
		return nil, err
	}

	d.db = open
	return open, nil
}

func (d *Database) Close() error {

	if d.db == nil {
		return errors.New("database is not initialized")
	}

	err := d.db.Close()

	d.db = nil

	return err
}
