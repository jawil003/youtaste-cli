package api

import (
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/services"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"net/http"
	"os"
)

func RegisterUser(api *gin.RouterGroup) {
	api.POST("/user/create", func(c *gin.Context) {
		var request models.CreateUserRequest

		err := c.BindJSON(&request)

		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		token, err := services.JWT().Create(models.Jwt{

			Firstname: request.Firstname, Lastname: request.Lastname,
		})

		if err != nil {
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

	api.GET("/user/me", func(context *gin.Context) {
		authorization, err := context.Cookie("token")

		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		if authorization == "" {
			context.JSON(400, gin.H{
				"error": "Authorization header is empty",
			})
			return
		}

		jwt := models.Jwt{}

		_, err = services.JWT().Decode(authorization, &jwt)

		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"firstname": jwt.Firstname,
			"lastname":  jwt.Lastname,
		})

	})

	api.GET("/admin", func(context *gin.Context) {

		clientIp := context.ClientIP()

		localAddr, err := services.Network().GetAddresses()

		isAdmin := funk.ContainsString(localAddr, clientIp)

		if err != nil {
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"isAdmin": isAdmin,
		})
	})
}
