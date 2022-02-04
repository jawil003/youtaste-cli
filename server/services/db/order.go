package db

import (
	"bs-to-scrapper/server/models"
	_ "github.com/thoas/go-funk"
	bolt "go.etcd.io/bbolt"
	"strings"
)

func (_ ServiceCollection) Order() OrderService {
	return OrderService{}
}

type OrderService struct {
}

func (o OrderService) GetOrdersByUser(user string) (*[]models.Order, error) {

	db, err := OpenDbConnection()

	if err != nil {
		return nil, err
	}

	var orderArray []models.Order

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("orders"))

		value := b.Get([]byte(user))

		var arrayRes []string

		arrayRes = strings.Split(string(value), ",")

		if len(arrayRes) == 0 {
			orderArray = []models.Order{}
			return nil
		}

		for _, orderJson := range arrayRes {

			if orderJson == "" {
				continue
			}

			order, err := models.ToOrder(orderJson)

			if err != nil {
				return err
			}

			orderArray = append(orderArray, order)
		}

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

	return &orderArray, nil

}

func (o OrderService) Create(orders []models.Order, user string) error {
	db, err := OpenDbConnection()

	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("orders"))

		value := b.Get([]byte(user))

		var jsonOrders []string

		for _, order := range orders {
			jsonOrder, err := order.ToJSON()

			if err != nil {
				return err
			}

			jsonOrders = append(jsonOrders, string(jsonOrder))
		}

		if len(value) > 0 {
			value := b.Get([]byte(user))

			arrayRes := strings.Split(string(value), ",")

			err := b.Put([]byte(user), []byte(strings.Join(append(arrayRes, jsonOrders...), ",")))

			if err != nil {
				return err
			}

			return nil
		} else {
			err := b.Put([]byte(user), []byte(strings.Join(jsonOrders, ",")))
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
