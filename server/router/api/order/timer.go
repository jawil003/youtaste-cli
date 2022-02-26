package order

import (
	"bs-to-scrapper/server/logger"
	"bs-to-scrapper/server/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterTimer(r *gin.RouterGroup, timerInst *services.TimerService) {

	timer := r.Group("/timer")

	timer.GET("", func(context *gin.Context) {

		res := gin.H{
			"time": timerInst.GetRemainingTime(),
		}

		logger.Logger().Info.Println(logger.LogResponse(http.StatusOK, res))
		context.JSON(http.StatusOK, res)
	})
}
