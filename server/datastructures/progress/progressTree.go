package progress

import "bs-to-scrapper/server/datastructures"

const (
	adminNew         = "ADMIN_NEW"
	chooseRestaurant = "CHOOSE_RESTAURANT"
	chooseMeals      = "CHOOSE_MEALS"
	done             = "DONE"
)

func ProgressTree() *datastructures.Tree {

	tree := &datastructures.Tree{Root: &datastructures.Node{Value: adminNew, Steps: []*datastructures.Node{{Value: chooseRestaurant, Steps: []*datastructures.Node{{Value: chooseMeals, Steps: []*datastructures.Node{{Value: done}}}}}}}}

	return tree
}
