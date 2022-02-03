package router

import (
	"bs-to-scrapper/server/services"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/api/orders", func(context *gin.Context) {

	})

	r.POST("/api/orders", func(context *gin.Context) {
		orders := context.GetStringSlice("orders")
		user := context.GetString("user")

		err := services.DB().Order().Create(orders, user)
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
