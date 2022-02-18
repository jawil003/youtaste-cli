package db

import bolt "go.etcd.io/bbolt"

func (_ ServiceCollection) Settings() SettingsService {
	return SettingsService{}
}

type SettingsService struct {
}

const (
	OrderTime          = "order_time"
	YoutastePhone      = "youtaste_phone"
	YoutastePassword   = "youtaste_password"
	LieferandoUsername = "lieferando_username"
	LieferandoPassword = "lieferando_password"
	ChoosenRestaurant  = "choosen_restaurant"
	RestaurantUrl      = "restaurant_url"
	OpeningTimes       = "opening_times"
)

func (_ SettingsService) Create(key string, value string) error {
	db, err := OpenDbConnection()
	defer db.Close()

	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("settings")).Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (_ SettingsService) Get(key string) (string, error) {
	db, err := OpenDbConnection()
	defer db.Close()

	if err != nil {
		return "", err
	}

	var value string

	err = db.View(func(tx *bolt.Tx) error {
		value = string(tx.Bucket([]byte("settings")).Get([]byte(key)))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return value, nil
}

func (_ SettingsService) ClearAll() error {
	db, err := OpenDbConnection()
	defer db.Close()

	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte("settings"))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
