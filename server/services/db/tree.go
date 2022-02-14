package db

import (
	"bs-to-scrapper/server/datastructures"
	"encoding/json"
	bolt "go.etcd.io/bbolt"
)

type TreeService struct {
}

func (_ ServiceCollection) Tree() *TreeService {
	return &TreeService{}
}

const (
	ProgressTree = "progress_tree"
)

func (_ TreeService) CreateOrUpdate(name string, tree datastructures.Tree) error {
	db, err := OpenDbConnection()
	defer db.Close()

	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("tree"))
		if err != nil {
			return err
		}

		res, err := json.Marshal(tree)

		if err != nil {
			return err
		}

		err = bucket.Put([]byte(name), res)
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

func (_ TreeService) Get(name string) (*datastructures.Tree, error) {
	db, err := OpenDbConnection()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	var tree datastructures.Tree

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("tree"))
		if bucket == nil {
			return nil
		}

		res := bucket.Get([]byte(name))
		if res == nil {
			return nil
		}

		err := json.Unmarshal(res, &tree)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &tree, nil
}
