package router

import (
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/services"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

type CreateRequest struct {
	User   string         `json:"user"`
	Orders []models.Order `json:"orders"`
}

type GetRequest struct {
	User string `form:"user"`
}

func Register(r *gin.Engine) {
	r.GET("/api/orders/:user", func(context *gin.Context) {

		user := context.Param("user")

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

	r.GET("/api/admin", func(context *gin.Context) {

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

	r.POST("/api/orders", func(context *gin.Context) {

		var request CreateRequest

		err := context.BindJSON(&request)

		if err != nil {
			context.JSON(500, gin.H{
				"error": err,
			})
		}

		err = services.DB().Order().Create(request.Orders, request.User)
		if err != nil {
			context.JSON(500, gin.H{
				"error": err,
			})
		}

		context.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.DELETE("/api/orders", func(context *gin.Context) {

	})

	r.PUT("/api/orders", func(context *gin.Context) {

	})
}
