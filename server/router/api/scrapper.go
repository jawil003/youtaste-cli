package api

import (
	"bs-to-scrapper/scrapper/youtaste"
	"github.com/gin-gonic/gin"
)

func RegisterScrapper(ap *gin.RouterGroup) {
	scrapper := ap.Group("/scrapper")
	{
		scrapper.GET("/openingTime", func(context *gin.Context) {
			page := youtaste.OpenInNewBrowserAndJoinYouTaste()

			page, err := youtaste.Login("+4917624615787", "HZWUKUGP42C9xG", page)

			if err != nil {
				context.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			page, err = youtaste.SearchForRestaurant("Restaurant am Markt", page)

			openingTimes, err := youtaste.GetOpeningTimes(page)
			if err != nil {
				context.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			if err != nil {
				context.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			context.JSON(200, gin.H{
				"openingtimes": openingTimes,
			})

		})

		scrapper.GET("/url", func(context *gin.Context) {
			page := youtaste.OpenInNewBrowserAndJoinYouTaste()

			page, err := youtaste.Login("", "", page)

			if err != nil {
				context.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			page, err = youtaste.SearchForRestaurant("Restaurant am Markt", page)

			if err != nil {
				context.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			url := youtaste.GetUrl(page)

			context.JSON(200, gin.H{
				"url": url,
			})

		})
	}
}
