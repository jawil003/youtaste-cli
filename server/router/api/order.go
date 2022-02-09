package api

import (
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RegisterOrders(api *gin.RouterGroup) {
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

		jwt := models.Jwt{}

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

	api.GET("/orders/user/:name", func(context *gin.Context) {

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

		jwt := models.Jwt{}

		_, err = services.JWT().Decode(token, &jwt)

		user := fmt.Sprintf("%s_%s", jwt.Firstname, jwt.Lastname)

		if user == "" {
			context.JSON(400, gin.H{
				"error": "user is required",
			})
			return
		}

		orders, err := services.DB().Order().GetByUser(user)

		for _, order := range *orders {
			if order.Name == context.Param("name") {
				context.JSON(200, gin.H{
					"order": order,
				})
				return
			}
		}

		context.JSON(200, gin.H{
			"error": "order not found",
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

		jwt := models.Jwt{}

		_, err = services.JWT().Decode(token, &jwt)

		var request models.CreateOrderRequest

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

	api.DELETE("/orders/user/:name", func(context *gin.Context) {
		token, err := context.Cookie("token")

		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		mealName := context.Param("name")

		customJWT := models.Jwt{}

		_, err = services.JWT().Decode(token, &customJWT)

		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		user := fmt.Sprintf("%s_%s", customJWT.Firstname, customJWT.Lastname)

		err = services.DB().Order().ClearByMealname(user, mealName)

		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		context.JSON(200, gin.H{"status": "ok"})
	})

	api.DELETE("/orders/user", func(context *gin.Context) {

		token, err := context.Cookie("token")

		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
		}

		custonJWT := models.Jwt{}

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
