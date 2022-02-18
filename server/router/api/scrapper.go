package api

import (
	"bs-to-scrapper/server/datastructures"
	"bs-to-scrapper/server/services/db"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
)

func RegisterScrapper(ap *gin.RouterGroup) {
	scrapper := ap.Group("/scrapper")
	{
		scrapper.GET("/openingTime", func(context *gin.Context) {

			res := os.Getenv(db.OpeningTimes)

			var weekdays *datastructures.Weekdays

			err := json.Unmarshal([]byte(res), &weekdays)
			if err != nil {
				context.JSON(400, gin.H{"error": err.Error()})
				return
			}

			context.JSON(200, gin.H{"weekdays": weekdays})

		})

		scrapper.GET("/url", func(context *gin.Context) {

			res := os.Getenv(db.RestaurantUrl)

			context.JSON(200, gin.H{"url": res})

		})
	}
}
