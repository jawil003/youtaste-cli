package models

type Poll struct {
	RestaurantName string `json:"restaurantName"`
	Provider       string `json:"provider"`
	Url            string `json:"url"`
}
