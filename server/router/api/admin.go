package api

import (
	"bs-to-scrapper/scrapper/lieferando"
	"bs-to-scrapper/scrapper/youtaste"
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
	ordertime, err := services.DB().Settings().Get(db.OrderTime)
	if err != nil {
		return
	}

	if ordertime != "" {
		_ = os.Setenv(db.OrderTime, ordertime)

	}

	youtastePhone, err := services.DB().Settings().Get(db.YoutastePhone)
	if err != nil {
		return
	}

	if youtastePhone != "" {
		_ = os.Setenv(db.YoutastePhone, youtastePhone)

	}

	youtastePassword, err := services.DB().Settings().Get(db.YoutastePassword)
	if err != nil {
		return
	}

	if youtastePassword != "" {
		_ = os.Setenv(db.YoutastePassword, youtastePassword)

	}

	lieferandoUsername, err := services.DB().Settings().Get(db.LieferandoUsername)
	if err != nil {
		return
	}

	if lieferandoUsername != "" {
		_ = os.Setenv(db.LieferandoUsername, lieferandoUsername)

	}

	lieferandoPassword, err := services.DB().Settings().Get(db.LieferandoPassword)
	if err != nil {
		return
	}

	if lieferandoPassword != "" {
		_ = os.Setenv(db.LieferandoPassword, lieferandoPassword)

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
		err = os.Setenv(db.OrderTime, createTimerRequest.OrderTime.Format(time.RFC3339))
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(db.OrderTime, createTimerRequest.OrderTime.String())
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv(db.YoutastePhone, createTimerRequest.YoutastePhone)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(db.YoutastePhone, createTimerRequest.YoutastePhone)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv(db.YoutastePassword, createTimerRequest.YoutastePassword)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(db.YoutastePassword, createTimerRequest.YoutastePassword)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv(db.LieferandoUsername, createTimerRequest.LieferandoUsername)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(db.LieferandoUsername, createTimerRequest.LieferandoUsername)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv(db.LieferandoPassword, createTimerRequest.LieferandoPassword)
		if err != nil {
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(db.LieferandoPassword, createTimerRequest.LieferandoPassword)
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

			switch tree.Tree.Root.Value {
			case progress.ChooseRestaurant:
				{
					highestPoll, err := services.DB().Poll().PersistFinalResult()

					if err != nil {
						context.JSON(400, gin.H{
							"error": err.Error(),
						})
						return
					}

					if highestPoll.Provider == "youtaste" {
						go services.Scrapper().ScrapUrlAndOpeningTimes(youtaste.YoutasteScrapper{}, *highestPoll)

					} else if highestPoll.Provider == "lieferando" {
						go services.Scrapper().ScrapUrlAndOpeningTimes(lieferando.LieferandoScrapper{}, *highestPoll)
					}

					break
				}

			}

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
