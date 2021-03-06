package api

import (
	"bs-to-scrapper/server/logger"
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/observer"
	"bs-to-scrapper/server/router/api/order"
	"bs-to-scrapper/server/router/api/poll"
	"bs-to-scrapper/server/services"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	api := r.Group("/api")

	api.Use(func(context *gin.Context) {

		if context.Request.URL.Path == "/api/user/create" || context.Request.URL.Path == "/api/polls/ws" || context.Request.URL.Path == "/api/progress/ws" || context.Request.URL.Path == "/api/admin/isAdmin" {
			logger.Logger().Info.Println("Skip authentication for ", context.Request.URL.Path)
			context.Next()
			return
		}

		authorization, err := context.Cookie("token")

		if err != nil {
			logger.Logger().Error.Println("Error while getting token from cookie: ", err)
			logger.Logger().Info.Println("Token is empty")
			context.JSON(404, gin.H{
				"error": "user doesn't exist",
			})
			context.Abort()
			return
		}

		if authorization == "" {
			logger.Logger().Info.Println("Token is empty")
			context.JSON(404, gin.H{
				"error": "user doesn't exist",
			})
			context.Abort()
			return
		}

		jwt := models.Jwt{}

		_, err = services.JWT().Decode(authorization, &jwt)

		if err != nil {
			logger.Logger().Error.Println(err)
			context.JSON(404, gin.H{
				"error": "user doesn't exist",
			})
			context.Abort()
			return
		}

		context.Set("user", jwt)
		context.Next()
	})

	timer := services.Timer()

	hub := observer.NewProgressObserverHub()

	RegisterUser(api)

	order.RegisterOrders(api, timer)

	poll.RegisterPolls(api)

	RegisterAdmin(api, timer, hub)

	RegisterProgress(api, hub)

	RegisterScrapper(api)

}
