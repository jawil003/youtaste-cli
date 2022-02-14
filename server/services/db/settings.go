package db

import bolt "go.etcd.io/bbolt"

func (_ ServiceCollection) Settings() SettingsService {
	return SettingsService{}
}

type SettingsService struct {
}

const (
	ORDER_TIME          = "order_time"
	YOUTASTE_PHONE      = "youtaste_phone"
	YOUTASTE_PASSWORD   = "youtaste_password"
	LIEFERANDO_USERNAME = "lieferando_username"
	LIEFERANDO_PASSWORD = "lieferando_password"
	POLL_ACTIVE         = "poll_active"
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
