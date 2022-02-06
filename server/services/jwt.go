package services

import (
	"github.com/golang-jwt/jwt"
	"os"
)

type JwtService struct {
}

func (_ JwtService) Create(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenString, err
}

func (_ JwtService) Decode(jwtToken string, claim jwt.Claims) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(jwtToken, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}
