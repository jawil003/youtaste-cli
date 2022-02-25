package api

import (
	"bs-to-scrapper/server/logger"
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func RegisterUser(api *gin.RouterGroup) {

	api.POST("/user/create", func(c *gin.Context) {
		var request models.CreateUserRequest

		err := c.BindJSON(&request)

		if err != nil {
			logger.Logger().Error.Println(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		token, err := services.JWT().Create(models.Jwt{

			Firstname: request.Firstname, Lastname: request.Lastname,
		})

		if err != nil {
			logger.Logger().Error.Println(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		ginMode := os.Getenv("GIN_MODE")

		if ginMode == "debug" {
			c.SetSameSite(http.SameSiteNoneMode)

			c.SetCookie("token", token, 60*60*24, "/", c.Request.Host, true, true)
		} else {
			c.SetCookie("token", token, 60*60*24, "/", c.Request.Host, false, true)
		}

		c.Status(200)

	})

	api.DELETE("/user", func(context *gin.Context) {

		context.SetCookie("token", "", -1, "/", context.Request.Host, false, true)

		context.Status(200)
	})

	api.GET("/user/me", func(context *gin.Context) {

		jwt, ok := context.Get("user")

		if !ok {
			logger.Logger().Error.Println("user not found")
			context.JSON(400, gin.H{
				"error": "user not found",
			})
			return
		}

		user := jwt.(models.Jwt)

		context.JSON(200, gin.H{
			"firstname": user.Firstname,
			"lastname":  user.Lastname,
		})

	})

}
