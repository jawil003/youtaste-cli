package api

import (
	"bs-to-scrapper/server/datastructures/progress"
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/observer"
	"bs-to-scrapper/server/services"
	"bs-to-scrapper/server/services/db"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"os"
	"time"
)

func initializeVariables(timerService *services.TimerService) {
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

	timeResolved, err := time.Parse(time.RFC3339, ordertime)
	if err != nil {
		return
	}

	timerService.Start(timeResolved)
}

func RegisterAdmin(r *gin.RouterGroup, timerService *services.TimerService, hub *observer.ProgressObserverHub) {

	initializeVariables(timerService)

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

		progressTree := services.DB().ProgressTree()

		if progressTree.Tree.Root.Value != progress.AdminNew {
			context.JSON(400, gin.H{
				"error": "progress tree is not in admin new state",
			})
			return
		}

		createTimerRequest := models.CreateAdminTimerRequest{}

		err := context.BindJSON(&createTimerRequest)
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
		err = os.Setenv(db.ORDER_TIME, createTimerRequest.OrderTime.Format(time.RFC3339))
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

		_, err = progressTree.Next(progressTree.Tree.Root.Steps[0].Value)

		timerService.Start(createTimerRequest.OrderTime)

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

	admin.PUT("/next", func(context *gin.Context) {

		tree := services.DB().ProgressTree()

		if tree.Tree.Root.Value == progress.AdminNew {
			context.JSON(400, gin.H{
				"error": "admin is only changeable by providing config",
			})
			return
		}

		if len(tree.Tree.Root.Steps) > 0 {
			next, err := tree.Next(tree.Tree.Root.Steps[0].Value)
			if err != nil {
				context.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			hub.SendAll(next.Root.Value)
			context.JSON(200, gin.H{
				"status": next.Root.Value,
			})
			return

		} else {
			context.JSON(400, gin.H{
				"error": "end state reached",
			})
			return
		}

	})
}
