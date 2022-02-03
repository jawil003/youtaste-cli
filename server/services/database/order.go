package database

import (
	"bs-to-scrapper/server/services"
	bolt "go.etcd.io/bbolt"
	"strings"
)

type Order struct {
}

func (o *Order) GetOrdersByUser(user string) (*[]string, error) {
	db, err := services.DB().Init()
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

	defer func(database *Database) {
		err := database.Close()
		if err != nil {
			panic(err)
		}
	}(services.DB())

	return &arrayRes, nil

}

func (o Order) Create(orders []string, user string) error {
	init, err := services.DB().Init()
	if err != nil {
		return err
	}

	err = init.Update(func(tx *bolt.Tx) error {
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
