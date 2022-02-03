package router

import (
	"bs-to-scrapper/server/models"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	User   string         `json:"user"`
	Orders []models.Order `json:"orders"`
}

func Register(r *gin.Engine) {
	r.GET("/api/orders", func(context *gin.Context) {

	})

	r.POST("/api/orders", func(context *gin.Context) {

		var request CreateRequest

		err := context.BindJSON(&request)

		if err != nil {
			context.JSON(500, gin.H{
				"error": err,
			})
		}

		/*err = services.DB().Order().Create(request.orders, request.user)
		if err != nil {
			context.JSON(500, gin.H{
				"error": err,
			})
		}

		context.JSON(200, gin.H{
			"message": "success",
		})*/
	})

	r.DELETE("/api/orders", func(context *gin.Context) {

	})

	r.PUT("/api/orders", func(context *gin.Context) {

	})
}
