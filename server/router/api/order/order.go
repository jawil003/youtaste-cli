package order

import (
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

func RegisterOrders(api *gin.RouterGroup) {

	pollTimerTime := os.Getenv("POLL_TIMER_TIME")
	pollTimerTimeAsInt, _ := strconv.Atoi(pollTimerTime)

	orderTimer := services.Timer()

	time.AfterFunc(time.Millisecond*time.Duration(pollTimerTimeAsInt), func() {
		orderTimerTime := os.Getenv("ORDER_TIMER_TIME")
		orderTimerTimeAsInt, err := strconv.Atoi(orderTimerTime)

		if err != nil {
			return
		}

		orderTimer.Start(int64(orderTimerTimeAsInt))
	})

	orderGroup := api.Group("/orders")

	RegisterTimer(orderGroup, orderTimer)

	orderGroup.GET("/user/:name", func(context *gin.Context) {

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

		var res models.Order

		for _, order := range *orders {
			if order.Name == context.Param("name") {
				res = order
				break
			}
		}

		if err != nil {
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"order": res,
		})
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
