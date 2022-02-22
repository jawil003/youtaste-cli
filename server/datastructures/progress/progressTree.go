package progress

import "bs-to-scrapper/server/datastructures"

const (
	AdminNew              = "ADMIN_NEW"
	ChooseRestaurant      = "CHOOSE_RESTAURANT"
	GetUrlAndOpeningTimes = "GET_URL_AND_OPENING_TIMES"
	ChooseMeals           = "CHOOSE_MEALS"
	Order                 = "ORDER"
	Done                  = "DONE"
)

func Tree() *datastructures.Tree {

	tree := &datastructures.Tree{Root: &datastructures.Node{Value: AdminNew, Steps: []*datastructures.Node{{Value: ChooseRestaurant, Steps: []*datastructures.Node{{Value: GetUrlAndOpeningTimes, Steps: []*datastructures.Node{{Value: ChooseMeals, Steps: []*datastructures.Node{{Value: Order, Steps: []*datastructures.Node{{Value: Done}}}}}}}}}}}}

	return tree
}
