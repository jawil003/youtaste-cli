package datastructures

import (
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

			return &tree, nil
		}
	}

	return nil, errors.New("option not found")
}
