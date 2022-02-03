package models

type Order struct {
	Name     string   `json:"name"`
	Variants []string `json:"variants"`
}
