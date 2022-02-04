package db

import (
	"bs-to-scrapper/server/models"
	"encoding/json"
	"errors"
	_ "github.com/thoas/go-funk"
	bolt "go.etcd.io/bbolt"
	"strings"
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

		valueString := string(value)

		var arrayRes []string

		err = json.Unmarshal([]byte(valueString), &arrayRes)

		if err != nil {
			return err
		}

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

			var arrayRes []string

			err := json.Unmarshal(value, &arrayRes)

			if err != nil {
				return err
			}

			jsonArray, err := json.Marshal(append(arrayRes, jsonOrders...))

			if err != nil {
				return err
			}

			err = b.Put([]byte(user), []byte(jsonArray))

			if err != nil {
				return err
			}

			return nil
		} else {

			jsonArray, err := json.Marshal(append(jsonOrders, jsonOrders...))

			if err != nil {
				return err
			}

			err = b.Put([]byte(user), []byte(jsonArray))
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

func (o OrderService) Update(orders []models.Order, user string) error {
	db, err := OpenDbConnection()

	if err != nil {
		CloseConnection(db)
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("orders"))

		value := string(b.Get([]byte(user)))

		if value == "" {
			return errors.New("no user found")
		}

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

			jsonArray, err := json.Marshal(append(arrayRes, jsonOrders...))

			if err != nil {
				return err
			}

			err = b.Put([]byte(user), []byte(jsonArray))

			if err != nil {
				return err
			}

			return nil
		} else {

			jsonArray, err := json.Marshal(append(jsonOrders, jsonOrders...))

			if err != nil {
				return err
			}

			err = b.Put([]byte(user), []byte(jsonArray))
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
