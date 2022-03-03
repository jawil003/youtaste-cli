package api

import (
	"bs-to-scrapper/server/datastructures"
	"bs-to-scrapper/server/datastructures/progress"
	"bs-to-scrapper/server/enums"
	"bs-to-scrapper/server/logger"
	"bs-to-scrapper/server/models"
	"bs-to-scrapper/server/observer"
	"bs-to-scrapper/server/scrapper/lieferando"
	"bs-to-scrapper/server/scrapper/youtaste"
	"bs-to-scrapper/server/services"
	"bs-to-scrapper/server/services/db"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"net/http"
	"os"
	"reflect"
	"time"
)

var infoLogger = logger.Logger().Info
var errorLogger = logger.Logger().Error
var warnLogger = logger.Logger().Warn

func initializeVariables(timerService *services.TimerService) {
	ordertime, err := services.DB().Settings().Get(enums.OrderTime)
	if err != nil {
		errorLogger.Panicln(err)
	}

	if ordertime != "" {
		_ = os.Setenv(enums.OrderTime, ordertime)

	}

	youtastePhone, err := services.DB().Settings().Get(enums.YoutastePhone)
	if err != nil {
		errorLogger.Panicln(err)
	}

	if youtastePhone != "" {
		_ = os.Setenv(enums.YoutastePhone, youtastePhone)

	}

	youtastePassword, err := services.DB().Settings().Get(enums.YoutastePassword)
	if err != nil {
		errorLogger.Panicln(err)
	}

	if youtastePassword != "" {
		_ = os.Setenv(enums.YoutastePassword, youtastePassword)

	}

	lieferandoUsername, err := services.DB().Settings().Get(enums.LieferandoUsername)
	if err != nil {
		errorLogger.Panicln(err)
	}

	if lieferandoUsername != "" {
		_ = os.Setenv(enums.LieferandoUsername, lieferandoUsername)

	}

	lieferandoPassword, err := services.DB().Settings().Get(enums.LieferandoPassword)
	if err != nil {
		errorLogger.Panic(err)
	}

	if lieferandoPassword != "" {
		_ = os.Setenv(enums.LieferandoPassword, lieferandoPassword)

	}

	if ordertime != "" {

		timeResolved, err := time.Parse(time.RFC3339, ordertime)
		if err != nil {
			errorLogger.Panic(err)
		}

		timerService.Start(timeResolved)
	}

	url, err := services.DB().Settings().Get(enums.RestaurantUrl)
	if err != nil {
		errorLogger.Panic(err)
	}

	if url != "" {
		_ = os.Setenv(enums.RestaurantUrl, url)
	}

	openingTimes, err := services.DB().Settings().Get(enums.OpeningTimes)
	if err != nil {
		errorLogger.Panicln(err)
	}
	if openingTimes != "" {
		_ = os.Setenv(enums.OpeningTimes, openingTimes)

		var openingTime *datastructures.Weekdays

		err = json.Unmarshal([]byte(openingTimes), &openingTime)
		if err != nil {
			errorLogger.Panicln(err)
		}

		currentWeekday := time.Now().Weekday().String()

		if err == nil {

			val := reflect.ValueOf(openingTime).Elem().FieldByName(currentWeekday).String()

			res, err := services.Time().ConvertOpeningTimeToDate(val)

			if err != nil {
				errorLogger.Panicln(err)
			}

			if res.After(time.Now()) {
				timerService.Start(*res)
			}

		}
	}
}

func RegisterAdmin(r *gin.RouterGroup, timerService *services.TimerService, hub *observer.ProgressObserverHub) {

	initializeVariables(timerService)

	admin := r.Group("/admin")

	isAdminHandler := func(context *gin.Context) {

		clientIp := context.ClientIP()

		localAddr, err := services.Network().GetAddresses()

		isAdmin := funk.ContainsString(localAddr, clientIp)

		if err != nil {
			errorLogger.Println(err)
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

		res := gin.H{
			"isAdmin": isAdmin,
		}

		logger.Logger().Info.Println(logger.LogResponse(http.StatusOK, res))
		context.JSON(200, res)
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
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = os.Setenv(enums.OrderTime, createTimerRequest.OrderTime.Format(time.RFC3339))
		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(enums.OrderTime, createTimerRequest.OrderTime.Format(time.RFC3339))
		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv(enums.YoutastePhone, createTimerRequest.YoutastePhone)
		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(enums.YoutastePhone, createTimerRequest.YoutastePhone)
		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv(enums.YoutastePassword, createTimerRequest.YoutastePassword)
		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(enums.YoutastePassword, createTimerRequest.YoutastePassword)
		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv(enums.LieferandoUsername, createTimerRequest.LieferandoUsername)
		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(enums.LieferandoUsername, createTimerRequest.LieferandoUsername)
		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = os.Setenv(enums.LieferandoPassword, createTimerRequest.LieferandoPassword)
		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = services.DB().Settings().Create(enums.LieferandoPassword, createTimerRequest.LieferandoPassword)
		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		next, err := progressTree.Next(progressTree.Tree.Root.Steps[0].Value)

		hub.SendAll(next.Root.Value)

		timerService.Start(createTimerRequest.OrderTime)

		if err != nil {
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		res := gin.H{
			"message": "success",
		}

		infoLogger.Println(logger.LogResponse(http.StatusOK, res))

		context.JSON(http.StatusOK, res)
	})

	admin.PUT("/next", func(context *gin.Context) {

		tree := services.DB().ProgressTree()

		if tree.Tree.Root.Value == progress.AdminNew {
			err := errors.New("admin is only changeable by providing config")
			errorLogger.Println(err)
			context.JSON(400, gin.H{
				"err": err.Error(),
			})
			return
		}

		if len(tree.Tree.Root.Steps) > 0 {

			switch tree.Tree.Root.Value {
			case progress.ChooseRestaurant:
				{
					highestPoll, err := services.DB().Poll().PersistFinalResult()

					if err != nil {
						errorLogger.Println(err)
						context.JSON(400, gin.H{
							"error": err.Error(),
						})
						return
					}

					if highestPoll.Provider == enums.YouTaste {
						go func() {
							defer func(tree *db.ProgressTreeService, option string) {
								next, _ := tree.Next(option)
								hub.SendAll(next.Root.Value)

							}(services.DB().ProgressTree(), services.DB().ProgressTree().Tree.Root.Steps[0].Value)

							res, err := services.Scrapper().ScrapUrlAndOpeningTimes(lieferando.Scrapper{}, *highestPoll)

							currentWeekday := time.Now().Weekday().String()

							if err == nil {

								val := reflect.ValueOf(res).Elem().FieldByName(currentWeekday).String()

								res, err := services.Time().ConvertOpeningTimeToDate(val)

								if err != nil {
									return
								}

								if res.After(time.Now()) {
									timerService.Start(*res)
								}

							} else {
								errorLogger.Println(err)
							}

						}()

					} else if highestPoll.Provider == enums.Lieferando {
						go func() {
							defer func() {

								progressTreeService := services.DB().ProgressTree()

								next, err := progressTreeService.Next(progressTreeService.Tree.Root.Steps[0].Value)
								if err != nil {
									errorLogger.Println(err)
									return
								}

								hub.SendAll(next.Root.Value)

							}()

							res, err := services.Scrapper().ScrapUrlAndOpeningTimes(lieferando.Scrapper{}, *highestPoll)

							currentWeekday := time.Now().Weekday().String()

							if err == nil {

								val := reflect.ValueOf(res).Elem().FieldByName(currentWeekday).String()

								res, err := services.Time().ConvertOpeningTimeToDate(val)

								if err != nil {
									errorLogger.Println(err)
									return
								}

								if res.After(time.Now()) {
									timerService.Start(*res)
								}

							} else {
								errorLogger.Println(err)
							}
						}()
					}

					break
				}
			case progress.GetUrlAndOpeningTimes:
				{
					warnLogger.Printf("need to wait for server action %s to complete", progress.GetUrlAndOpeningTimes)
					context.JSON(400, gin.H{"error": "need to wait for server action to complete"})
					return
				}
			case progress.ChooseMeals:
				{

					highestPoll, err := services.DB().Poll().PersistFinalResult()

					jwt, ok := context.Get("user")

					if !ok {
						errorLogger.Println("user not found")
						context.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
					}

					user := jwt.(models.Jwt)

					username := services.User().GetUsername(user.Firstname, user.Lastname)

					orders, err := services.DB().Order().GetByUser(username)

					if err != nil {
						errorLogger.Println(err)
						context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
					}

					if highestPoll.Provider == enums.YouTaste {
						go func() {
							err = services.Scrapper().OrderMeals(youtaste.Scrapper{}, *highestPoll, *orders)
							if err != nil {
								return
							}
						}()
					} else {
						go func() {
							err = services.Scrapper().OrderMeals(lieferando.Scrapper{}, *highestPoll, *orders)
							if err != nil {
								return
							}
						}()
					}
				}
			case progress.Order:
				{
					warnLogger.Printf("need to wait for server action %s to complete", progress.Order)
					context.JSON(400, gin.H{"error": "need to wait for server action to complete"})
					return
				}
			}

			next, err := tree.Next(tree.Tree.Root.Steps[0].Value)
			if err != nil {
				errorLogger.Println(err)
				context.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}

			hub.SendAll(next.Root.Value)
			infoLogger.Printf("Broadcast client status %s", next.Root.Value)

			res := gin.H{
				"status": next.Root.Value,
			}

			infoLogger.Println(logger.LogResponse(http.StatusOK, res))
			context.JSON(200, res)
			return

		} else {

			res := gin.H{
				"error": "end state reached",
			}

			errorLogger.Println(logger.LogResponse(http.StatusBadRequest, res))
			context.JSON(http.StatusBadRequest, res)
			return
		}

	})
}
