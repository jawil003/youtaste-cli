package router

import (
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/thoas/go-funk"
	"net/http"
	"os"
)

type CreateRequest struct {
	Orders []models.Order `json:"orders"`
}

type CreateUserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type CustonJWT struct {
	jwt.StandardClaims
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type GetRequest struct {
	User string `form:"user"`
}

func Register(r *gin.Engine) {

	api := r.Group("/api")

	api.POST("/user/create", func(c *gin.Context) {
		var request CreateUserRequest

		err := c.BindJSON(&request)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		token, err := services.JWT().Create(CustonJWT{

			Firstname: request.Firstname, Lastname: request.Lastname,
		})

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		ginMode := os.Getenv("GIN_MODE")

		if ginMode == "debug" {
			c.SetSameSite(http.SameSiteNoneMode)

			c.SetCookie("token", token, 60, "/", c.Request.Host, true, true)
		} else {
			c.SetCookie("token", token, 60, "/", c.Request.Host, false, true)
		}

		c.Status(200)

	})

	api.GET("/user/me", func(context *gin.Context) {
		authorization, err := context.Cookie("token")

		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		if authorization == "" {
			context.JSON(400, gin.H{
				"error": "Authorization header is empty",
			})
			return
		}

		jwt := CustonJWT{}

		_, err = services.JWT().Decode(authorization, &jwt)

		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"firstname": jwt.Firstname,
			"lastname":  jwt.Lastname,
		})

	})

	api.GET("/orders/user", func(context *gin.Context) {

		token, err := context.Cookie("token")

		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		if token == "" {
			context.JSON(400, gin.H{
				"error": "Authorization header is empty",
			})
			return
		}

		jwt := CustonJWT{}

		_, err = services.JWT().Decode(token, &jwt)

		user := fmt.Sprintf("%s_%s", jwt.Firstname, jwt.Lastname)

		if user == "" {
			context.JSON(400, gin.H{
				"error": "user is required",
			})
			return
		}

		orders, err := services.DB().Order().GetByUser(user)

		if err != nil {
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"orders": orders,
		})
	})

	api.GET("/admin", func(context *gin.Context) {

		clientIp := context.ClientIP()

		localAddr, err := services.Network().GetAddresses()

		isAdmin := funk.ContainsString(localAddr, clientIp)

		if err != nil {
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"isAdmin": isAdmin,
		})
	})

	api.POST("/orders", func(context *gin.Context) {

		token, err := context.Cookie("token")

		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		jwt := CustonJWT{}

		_, err = services.JWT().Decode(token, &jwt)

		var request CreateRequest

		err = context.BindJSON(&request)

		if err != nil {
			context.JSON(500, gin.H{
				"error": err,
			})
		}

		if err != nil {
			context.JSON(500, gin.H{
				"error": err,
			})
		}

		err = services.DB().Order().Create(request.Orders, fmt.Sprintf("%s_%s", jwt.Firstname, jwt.Lastname))
		if err != nil {
			context.JSON(500, gin.H{
				"error": err,
			})
		}

		context.JSON(200, gin.H{
			"message": "success",
		})
	})

	api.DELETE("/orders/user", func(context *gin.Context) {

		token, err := context.Cookie("token")

		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
		}

		custonJWT := CustonJWT{}

		_, err = services.JWT().Decode(token, &custonJWT)

		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
		}

		user := fmt.Sprintf("%s_%s", custonJWT.Firstname, custonJWT.Lastname)

		err = services.DB().Order().Clear(user)

		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"status": "ok",
		})
	})

	api.DELETE("/orders/all", func(context *gin.Context) {
		err := services.DB().Order().ClearAll()
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"status": "ok",
		})
	})

}
