package ws

import (
	"bs-to-scrapper/server/router/api"
	"bs-to-scrapper/server/services"
	"github.com/gin-gonic/gin"
)

type Poll struct {
	RestaurantName string `json:"restaurantName"`
}

func RegisterPolls(r *gin.RouterGroup) {

	pollsGroup := r.Group("/polls")

	pollsGroup.POST("/new", func(context *gin.Context) {
		token, err := context.Cookie("token")

		if err != nil {
			context.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		if token == "" {
			context.JSON(400, gin.H{
				"error": "Authorization header is empty",
			})
			return
		}

		jwt := api.CustomJWT{}

		_, err = services.JWT().Decode(token, &jwt)

		var poll Poll

		err = context.BindJSON(&poll)

		if err != nil {
			context.JSON(400, gin.H{"error": err.Error()})
			return
		}
	})

	pollsGroup.POST("/choose", func(context *gin.Context) {

	})
}
