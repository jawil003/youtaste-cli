package models

import "github.com/golang-jwt/jwt"

type Jwt struct {
	jwt.StandardClaims
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
