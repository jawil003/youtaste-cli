package datastructures

import (
	"bs-to-scrapper/server/services"
	"bs-to-scrapper/server/services/db"
	"errors"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Value string
	Steps []*Node
}

func (tree *Tree) Next(Option string) (*Tree, error) {
	options := tree.Root.Steps

	for _, option := range options {
		if option.Value == Option {
			tree := Tree{Root: option}

			err := services.DB().Tree().CreateOrUpdate(db.ProgressTree, tree)

			if err != nil {
				return nil, err
			}

			return &tree, nil
		}
	}

	return nil, errors.New("option not found")
}

func NewProgressTree() (*Tree, error) {
	tree := Tree{Root: &Node{Value: "ADMIN_NEW", Steps: []*Node{{Value: "CHOOSE_RESTAURANT", Steps: []*Node{{Value: "CHOOSE_MEALS", Steps: []*Node{{Value: "DONE"}}}}}}}}

	err := services.DB().Tree().CreateOrUpdate(db.ProgressTree, tree)

	if err != nil {
		return nil, err
	}

	return &tree, nil
}
