package api

import (
	"bs-to-scrapper/server/models"
	socket_io "bs-to-scrapper/server/router/api/socket.io"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type CreateRequest struct {
	Orders []models.Order `json:"orders"`
}

type CreateUserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type CustomJWT struct {
	jwt.StandardClaims
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type GetRequest struct {
	User string `form:"user"`
}

func Register(r *gin.Engine) {

	api := r.Group("/api")

	RegisterUser(api)

	RegisterOrders(api)

	socket_io.RegisterSocketIO(api)

}
