package models

type CreateOrderRequest struct {
	Orders []Order `json:"orders"`
}
