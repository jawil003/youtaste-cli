package api

import (
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/services"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"os"
)

func RegisterAdmin(r *gin.RouterGroup) {

	admin := r.Group("/admin")

	isAdminHandler := func(context *gin.Context) {

		clientIp := context.ClientIP()

		localAddr, err := services.Network().GetAddresses()

		isAdmin := funk.ContainsString(localAddr, clientIp)

		if err != nil {
			context.Abort()
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.Set("isAdmin", isAdmin)

		context.Next()
	}

	admin.Use(isAdminHandler)

	admin.GET("/isAdmin", func(context *gin.Context) {
		isAdminHandler(context)

		isAdmin := context.GetBool("isAdmin")

		context.JSON(200, gin.H{
			"isAdmin": isAdmin,
		})
	})

	admin.POST("/timer", func(context *gin.Context) {

		var createTimerRequest *models.CreateAdminTimerRequest

		err := context.BindJSON(createTimerRequest)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv("POLL_TIMER_TIME", string(rune(createTimerRequest.PollTime)))
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = os.Setenv("ORDER_TIMER_TIME", string(rune(createTimerRequest.OrderTime)))
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"message": "success",
		})
	})
}
