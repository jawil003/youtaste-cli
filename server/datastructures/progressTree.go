package datastructures

const (
	adminNew         = "ADMIN_NEW"
	chooseRestaurant = "CHOOSE_RESTAURANT"
	chooseMeals      = "CHOOSE_MEALS"
	done             = "DONE"
)

func ProgressTree() *Tree {

	tree := &Tree{Root: &Node{Value: adminNew, Steps: []*Node{{Value: chooseRestaurant, Steps: []*Node{{Value: chooseMeals, Steps: []*Node{{Value: done}}}}}}}}

	return tree
}
