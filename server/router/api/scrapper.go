package api

import (
	"bs-to-scrapper/scrapper/youtaste"
	"bs-to-scrapper/server/services/db"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func RegisterScrapper(ap *gin.RouterGroup) {
	scrapper := ap.Group("/scrapper")
	{
		scrapper.GET("/openingTime", func(context *gin.Context) {
			page, err := youtaste.OpenInNewBrowserAndJoinYouTaste()

			if err != nil {
				context.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			page = page.Timeout(2 * time.Second)

			page, err = youtaste.Login(os.Getenv(db.YOUTASTE_PHONE), os.Getenv(db.YOUTASTE_PASSWORD), page)

			if err != nil {

				switch err.(type) {
				default:
					context.JSON(400, gin.H{
						"error": err.Error(),
					})
					break

				}

				return
			}

			page, err = youtaste.SearchForRestaurant("Restaurant am Markt", page)

			openingTimes, err := youtaste.GetOpeningTimes(page)
			if err != nil {
				context.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			if err != nil {
				context.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			context.JSON(200, gin.H{
				"openingtimes": openingTimes,
			})

		})

		scrapper.GET("/url", func(context *gin.Context) {
			page, err := youtaste.OpenInNewBrowserAndJoinYouTaste()

			if err != nil {
				context.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			page, err = youtaste.Login(os.Getenv(db.YOUTASTE_PHONE), os.Getenv(db.YOUTASTE_PASSWORD), page)

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

			url, err := youtaste.GetUrl(page)

			if err != nil {
				context.JSON(500, gin.H{
					"error": err.Error(),
				})
				return
			}

			context.JSON(200, gin.H{
				"url": url,
			})

		})
	}
}
