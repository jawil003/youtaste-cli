package poll

import (
	"bs-to-scrapper/server/services"
	"github.com/gin-gonic/gin"
)

func RegisterTimer(r *gin.RouterGroup) {
	timer := r.Group("/timer")

	timer.GET("", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"time": services.Timer().GetRemainingTime(),
		})
	})
}
