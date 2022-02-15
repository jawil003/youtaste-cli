package db

import (
	"bs-to-scrapper/server/datastructures"
	"bs-to-scrapper/server/datastructures/progress"
	"encoding/json"
	bolt "go.etcd.io/bbolt"
)

type ProgressTreeService struct {
	Tree *datastructures.Tree
}

var treeService *ProgressTreeService

func (_ ServiceCollection) ProgressTree() *ProgressTreeService {

	if treeService == nil {

		tree := progress.ProgressTree()

		treeFromDb, err := getFromDb(ProgressTree)

		if err != nil {
			return nil
		}

		if treeFromDb.Root != nil {
			tree = treeFromDb
		} else {
			err := createOrUpdate(ProgressTree, *tree)
			if err != nil {
				return nil
			}
		}

		treeService = &ProgressTreeService{
			Tree: tree,
		}
	}

	return treeService
}

const (
	ProgressTree = "progress_tree"
)

func (ts *ProgressTreeService) Next(option string) (*datastructures.Tree, error) {
	next, err := ts.Tree.Next(option)
	if err != nil {
		return nil, err
	}

	err = createOrUpdate(ProgressTree, *next)
	if err != nil {
		return nil, err
	}

	ts.Tree = next

	return next, nil
}

func (ts *ProgressTreeService) Reset() (*datastructures.Tree, error) {
	err := clear(ProgressTree)

	if err != nil {
		return nil, err
	}

	tree := progress.ProgressTree()

	ts.Tree = tree

	return tree, nil
}

func createOrUpdate(name string, tree datastructures.Tree) error {
	db, err := OpenDbConnection()
	defer db.Close()

	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("ProgressTree"))
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

func getFromDb(name string) (*datastructures.Tree, error) {
	db, err := OpenDbConnection()
	defer db.Close()

	if err != nil {
		return nil, err
	}

	var tree datastructures.Tree

	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("ProgressTree"))
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

func clear(name string) error {
	db, err := OpenDbConnection()
	defer db.Close()

	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("ProgressTree"))
		if bucket == nil {
			return nil
		}

		err := bucket.Delete([]byte(name))
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
