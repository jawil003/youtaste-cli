package models

import "encoding/json"

type Order struct {
	Name     string   `json:"name"`
	Variants []string `json:"variants"`
}

func (order Order) ToJSON() ([]byte, error) {
	return json.Marshal(order)
}

func ToOrder(jsonVal string) (Order, error) {

	var order Order

	err := json.Unmarshal([]byte(jsonVal), &order)

	return order, err
}
