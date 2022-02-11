package models

type PollWithCount struct {
	Poll
	Count int `json:"count"`
}
