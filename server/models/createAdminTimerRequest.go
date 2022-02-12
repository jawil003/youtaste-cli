package models

type CreateAdminTimerRequest struct {
	PollTime  int `json:"pollTime"`
	OrderTime int `json:"orderTime"`
}
