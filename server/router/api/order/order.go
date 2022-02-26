package order

import (
	"bs-to-scrapper/server/logger"
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var infoLogger = logger.Logger().Info
var errorLogger = logger.Logger().Error

func RegisterOrders(api *gin.RouterGroup, orderTimer *services.TimerService) {

	orderGroup := api.Group("/orders")

	RegisterTimer(orderGroup, orderTimer)

	orderGroup.GET("/user/:name", func(context *gin.Context) {

		jwt, ok := context.Get("user")

		if !ok {
			infoLogger.Println("JWT not found")
			context.JSON(401, gin.H{"error": "unauthorized"})
			return
		}

		user := services.User().GetUsername(jwt.(models.Jwt).Firstname, jwt.(models.Jwt).Lastname)

		if user == "" {
			errorLogger.Println("JWT not found")
			context.JSON(400, gin.H{
				"error": "user is required",
			})
			return
		}

		orders, err := services.DB().Order().GetByUser(user)

		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		var order models.Order

		for _, locOrder := range *orders {
			if order.Name == context.Param("name") {
				order = locOrder
				break
			}
		}

		res := gin.H{
			"order": order,
		}

		infoLogger.Println(logger.LogResponse(http.StatusOK, res))

		context.JSON(http.StatusOK, res)
	})

	orderGroup.GET("/user", func(context *gin.Context) {

		jwt, ok := context.Get("user")

		if !ok {
			context.JSON(401, gin.H{"error": "unauthorized"})
			return
		}

		user := services.User().GetUsername(jwt.(models.Jwt).Firstname, jwt.(models.Jwt).Lastname)

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

	orderGroup.POST("", func(context *gin.Context) {

		jwt, ok := context.Get("user")

		if !ok {
			context.JSON(401, gin.H{"error": "unauthorized"})
			return
		}

		var request models.CreateOrderRequest

		err := context.BindJSON(&request)

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

		err = services.DB().Order().Create(request.Orders, services.User().GetUsername(jwt.(models.Jwt).Firstname, jwt.(models.Jwt).Lastname))
		if err != nil {
			context.JSON(500, gin.H{
				"error": err,
			})
		}

		context.JSON(200, gin.H{
			"message": "success",
		})
	})

	orderGroup.DELETE("/user/:name", func(context *gin.Context) {
		mealName := context.Param("name")

		customJWT, ok := context.Get("user")

		if !ok {
			context.JSON(401, gin.H{"error": "unauthorized"})
			return
		}

		user := fmt.Sprintf(services.User().GetUsername(customJWT.(models.Jwt).Firstname, customJWT.(models.Jwt).Lastname))

		err := services.DB().Order().ClearByMealname(user, mealName)

		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}

		context.JSON(200, gin.H{"status": "ok"})
	})

	orderGroup.DELETE("/user", func(context *gin.Context) {

		custonJWT, ok := context.Get("user")

		if !ok {
			context.JSON(401, gin.H{"error": "unauthorized"})
			return
		}

		user := fmt.Sprintf(services.User().GetUsername(custonJWT.(models.Jwt).Firstname, custonJWT.(models.Jwt).Lastname))

		err := services.DB().Order().Clear(user)

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

	orderGroup.DELETE("/all", func(context *gin.Context) {
		err := services.DB().Order().ClearAll()
		if err != nil {
			errorLogger.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		res := gin.H{
			"status": "ok",
		}

		infoLogger.Println(logger.LogResponse(http.StatusOK, res))
		context.JSON(http.StatusOK, res)
	})
}
