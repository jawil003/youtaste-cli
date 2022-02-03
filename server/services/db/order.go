package db

import (
	bolt "go.etcd.io/bbolt"
	"strings"
)

func (_ ServiceCollection) Order() OrderService {
	return OrderService{}
}

type OrderService struct {
}

func (o *OrderService) GetOrdersByUser(user string) (*[]string, error) {

	db, err := OpenDbConnection()

	if err != nil {
		return nil, err
	}

	var arrayRes []string

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("orders"))

		value := b.Get([]byte(user))

		arrayRes = strings.Split(string(value), ",")

		return nil
	})

	if err != nil {
		return nil, err
	}

	defer func(db *bolt.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	return &arrayRes, nil

}

func (o OrderService) Create(orders []string, user string) error {
	db, err := OpenDbConnection()

	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("orders"))

		value := b.Get([]byte(user))

		if len(value) > 0 {
			value := b.Get([]byte(user))

			arrayRes := strings.Split(string(value), ",")

			err := b.Put([]byte(user), []byte(strings.Join(append(arrayRes, orders...), ",")))

			if err != nil {
				return err
			}

			return nil
		} else {
			err := b.Put([]byte(user), []byte(strings.Join(orders, ",")))
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
