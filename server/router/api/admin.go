package api

import (
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/services"
	"bs-to-scrapper/server/services/db"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"os"
)

func initializeVariables() {
	ordertime, err := services.DB().Settings().Get(db.ORDER_TIME)
	if err != nil {
		return
	}

	if ordertime != "" {
		_ = os.Setenv(db.ORDER_TIME, ordertime)

	}

	youtastePhone, err := services.DB().Settings().Get(db.YOUTASTE_PHONE)
	if err != nil {
		return
	}

	if youtastePhone != "" {
		_ = os.Setenv(db.YOUTASTE_PHONE, youtastePhone)

	}

	youtastePassword, err := services.DB().Settings().Get(db.YOUTASTE_PASSWORD)
	if err != nil {
		return
	}

	if youtastePassword != "" {
		_ = os.Setenv(db.YOUTASTE_PASSWORD, youtastePassword)

	}

	lieferandoUsername, err := services.DB().Settings().Get(db.LIEFERANDO_USERNAME)
	if err != nil {
		return
	}

	if lieferandoUsername != "" {
		_ = os.Setenv(db.LIEFERANDO_USERNAME, lieferandoUsername)

	}

	lieferandoPassword, err := services.DB().Settings().Get(db.LIEFERANDO_PASSWORD)
	if err != nil {
		return
	}

	if lieferandoPassword != "" {
		_ = os.Setenv(db.LIEFERANDO_PASSWORD, lieferandoPassword)

	}
}

func RegisterAdmin(r *gin.RouterGroup) {

	initializeVariables()

	admin := r.Group("/admin")

	isAdminHandler := func(context *gin.Context) {

		clientIp := context.ClientIP()

		localAddr, err := services.Network().GetAddresses()

		isAdmin := funk.ContainsString(localAddr, clientIp)

		if err != nil {
			context.Abort()
			context.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.Set("isAdmin", isAdmin)

		context.Next()
	}

	admin.Use(isAdminHandler)

	admin.GET("/isAdmin", func(context *gin.Context) {
		isAdminHandler(context)

		isAdmin := context.GetBool("isAdmin")

		context.JSON(200, gin.H{
			"isAdmin": isAdmin,
		})
	})

	admin.POST("/set", func(context *gin.Context) {

		var createTimerRequest *models.CreateAdminTimerRequest

		err := context.BindJSON(createTimerRequest)
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
		err = os.Setenv(db.ORDER_TIME, createTimerRequest.OrderTime.String())
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(db.ORDER_TIME, createTimerRequest.OrderTime.String())
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv(db.YOUTASTE_PHONE, createTimerRequest.YoutastePhone)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(db.YOUTASTE_PHONE, createTimerRequest.YoutastePhone)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv(db.YOUTASTE_PASSWORD, createTimerRequest.YoutastePassword)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(db.YOUTASTE_PASSWORD, createTimerRequest.YoutastePassword)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv(db.LIEFERANDO_USERNAME, createTimerRequest.LieferandoUsername)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(db.LIEFERANDO_USERNAME, createTimerRequest.LieferandoUsername)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv(db.LIEFERANDO_PASSWORD, createTimerRequest.LieferandoPassword)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(db.LIEFERANDO_PASSWORD, createTimerRequest.LieferandoPassword)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"message": "success",
		})
	})

	admin.POST("/lieferando", func(context *gin.Context) {

		var createLieferandoRequest *models.CreateProviderLoginRequest

		err := context.BindJSON(createLieferandoRequest)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv("LIEFERANDO_USERNAME", createLieferandoRequest.Username)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = os.Setenv("LIEFERANDO_PASSWORD", createLieferandoRequest.Password)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"message": "success",
		})
	})

	admin.POST("/youtaste", func(context *gin.Context) {
		var createLieferandoRequest *models.CreateProviderLoginRequest

		err := context.BindJSON(createLieferandoRequest)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv("YOUTASTE_PHONE", createLieferandoRequest.Phone)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = os.Setenv("YOUTASTE_PASSWORD", createLieferandoRequest.Password)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		context.JSON(200, gin.H{
			"message": "success",
		})
	})
}
