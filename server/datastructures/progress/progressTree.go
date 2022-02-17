package progress

import "bs-to-scrapper/server/datastructures"

const (
	AdminNew         = "ADMIN_NEW"
	ChooseRestaurant = "CHOOSE_RESTAURANT"
	ChooseMeals      = "CHOOSE_MEALS"
	Done             = "DONE"
)

func ProgressTree() *datastructures.Tree {

	tree := &datastructures.Tree{Root: &datastructures.Node{Value: AdminNew, Steps: []*datastructures.Node{{Value: ChooseRestaurant, Steps: []*datastructures.Node{{Value: ChooseMeals, Steps: []*datastructures.Node{{Value: Done}}}}}}}}

	return tree
}
