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

			arrayResNew := append(arrayRes)
			ordersNew := append(orders)

			for ordersElemIndex, ordersElem := range orders {

				var arrayResItemFound *models.Order
				var arrayResItemFoundIndex *int
				for arrayResItemIndex, arrayResItem := range arrayRes {
					if arrayResItem.Name == ordersElem.Name {
						arrayResItemFound = &arrayResItem
						arrayResItemFoundIndex = &arrayResItemIndex
					}
				}

				if arrayResItemFound != nil {
					ordersNew = append(ordersNew[:ordersElemIndex], ordersNew[ordersElemIndex+1:]...)
					arrayResNew = append(arrayResNew[:*arrayResItemFoundIndex], arrayResNew[*arrayResItemFoundIndex+1:]...)

					arrayResItemFound = &ordersElem

					arrayResNew = append(arrayResNew, *arrayResItemFound)

					jsonArray, err := json.Marshal(arrayResNew)

					if err != nil {
						return err
					}

					err = b.Put([]byte(user), jsonArray)

					return nil
				}
			}

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

func (o OrderService) Clear(user string) error {
	db, err := OpenDbConnection()

	if err != nil {
		CloseConnection(db)
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("orders"))

		err := b.Delete([]byte(user))
		if err != nil {
			return err
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

func (o OrderService) ClearAll() error {
	db, err := OpenDbConnection()

	if err != nil {
		CloseConnection(db)
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte("orders"))

		if err != nil {
			return err
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
