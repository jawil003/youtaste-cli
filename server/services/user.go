package services

import "strings"

type UserService struct {
}

func (_ UserService) GetUsername(firstname string, lastname string) string {
	return strings.ToLower(firstname) + " " + strings.ToLower(lastname)
}
