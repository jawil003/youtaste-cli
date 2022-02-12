package api

import (
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/router/api/order"
	"bs-to-scrapper/server/router/api/poll"
	"bs-to-scrapper/server/services"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {

	api := r.Group("/api")

	api.Use(func(context *gin.Context) {

		if context.Request.URL.Path == "/api/user/create" || context.Request.URL.Path == "/api/polls/ws" {
			context.Next()
			return
		}

		authorization, err := context.Cookie("token")

		if err != nil {
			context.JSON(404, gin.H{
				"error": "user doesn't exist",
			})
			context.Abort()
			return
		}

		if authorization == "" {
			context.JSON(404, gin.H{
				"error": "user doesn't exist",
			})
			context.Abort()
			return
		}

		jwt := models.Jwt{}

		_, err = services.JWT().Decode(authorization, &jwt)

		if err != nil {
			context.JSON(404, gin.H{
				"error": "user doesn't exist",
			})
			context.Abort()
			return
		}

		context.Set("user", jwt)
		context.Next()
	})

	pollTimer := services.Timer()

	RegisterUser(api, pollTimer)

	order.RegisterOrders(api)

	poll.RegisterPolls(api, pollTimer)

	RegisterAdmin(api)

}
