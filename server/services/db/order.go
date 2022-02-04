package db

import (
	"bs-to-scrapper/server/models"
	"encoding/json"
	_ "github.com/thoas/go-funk"
	bolt "go.etcd.io/bbolt"
)

func (_ ServiceCollection) Order() OrderService {
	return OrderService{}
}

type OrderService struct {
}

func (o OrderService) GetByUser(user string) (*[]models.Order, error) {

	db, err := OpenDbConnection()

	if err != nil {
		CloseConnection(db)
		return nil, err
	}

	var orderArray []models.Order

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("orders"))

		value := b.Get([]byte(user))

		if value == nil {
			orderArray = []models.Order{}
			return nil
		}

		err = json.Unmarshal(value, &orderArray)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		CloseConnection(db)
		return nil, err
	}

	CloseConnection(db)

	return &orderArray, nil

}

func (o OrderService) Create(orders []models.Order, user string) error {
	db, err := OpenDbConnection()

	if err != nil {
		CloseConnection(db)
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("orders"))

		value := b.Get([]byte(user))

		var arrayRes []models.Order

		if value != nil {
			err := json.Unmarshal(value, &arrayRes)

			if err != nil {
				return err
			}
		}

		if len(arrayRes) > 0 {

			jsonArray, err := json.Marshal(append(arrayRes, orders...))

			if err != nil {
				return err
			}

			err = b.Put([]byte(user), jsonArray)

			if err != nil {
				return err
			}

			return nil
		} else {

			jsonArray, err := json.Marshal(orders)

			if err != nil {
				return err
			}

			err = b.Put([]byte(user), jsonArray)
			if err != nil {
				return err
			}
		}

		return nil
	})

	CloseConnection(db)

	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}
