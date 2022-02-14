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

var progressTree *Tree

func ResetProgressTree() (*Tree, error) {
	progressTree = nil

	err := services.DB().Tree().Clear(db.ProgressTree)
	if err != nil {
		return nil, err
	}

	tree, err := NewProgressTree()

	if err != nil {
		return nil, err
	}

	return tree, nil
}

func NewProgressTree() (*Tree, error) {

	if progressTree != nil {
		return progressTree, nil
	}

	tree, err := services.DB().Tree().Get(db.ProgressTree)

	if err != nil {
		return nil, err
	}

	if tree != nil {
		return tree, nil
	}

	tree = &Tree{Root: &Node{Value: "ADMIN_NEW", Steps: []*Node{{Value: "CHOOSE_RESTAURANT", Steps: []*Node{{Value: "CHOOSE_MEALS", Steps: []*Node{{Value: "DONE"}}}}}}}}

	err = services.DB().Tree().CreateOrUpdate(db.ProgressTree, *tree)

	if err != nil {
		return nil, err
	}

	progressTree = tree

	return tree, nil
}
