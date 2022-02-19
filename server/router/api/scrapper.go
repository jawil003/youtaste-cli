package api

import (
	"bs-to-scrapper/server/datastructures"
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/services"
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

			restaurant, err := services.DB().Settings().Get(db.ChoosenRestaurant)

			if err != nil {
				context.JSON(400, gin.H{"error": err.Error()})
				return
			}

			var highestPoll *models.PollWithCount

			err = json.Unmarshal([]byte(restaurant), &highestPoll)
			if err != nil {
				context.JSON(400, gin.H{"error": err.Error()})
				return
			}

			res := os.Getenv(db.RestaurantUrl)

			if res == "" {

				switch highestPoll.Provider {
				case "youtaste":
					res = "https://www.youtaste.com"
				case "lieferando":
					res = "https://www.lieferando.de"
				}

				context.JSON(200, gin.H{"url": res, "pending": true, "provider": highestPoll.Provider})
				return

			}

			context.JSON(200, gin.H{"url": res, "pending": false, "provider": highestPoll.Provider})

		})
	}
}
