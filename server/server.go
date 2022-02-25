package server

import (
	"bs-to-scrapper/server/enums"
	"bs-to-scrapper/server/logger"
	"bs-to-scrapper/server/router/api"
	"bs-to-scrapper/server/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/jweslley/localtunnel"
	"github.com/thoas/go-funk"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type OrderWithRes struct {
	Restaurant string
	Order      []Product
}

type Product struct {
	Name    string
	Variant []string
}

func Serve() {

	f, err := os.OpenFile(enums.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Logger().Error.Panic(err)
	}

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	err = godotenv.Load()
	if err != nil {
		logger.Logger().Error.Panic(err)
	}

	r := gin.Default()
	err = r.SetTrustedProxies(nil)
	if err != nil {
		logger.Logger().Error.Panic(err)
	}

	if os.Getenv("GIN_MODE") == "debug" {

		corsUrlConn := os.Getenv("CORS_URL")

		corsUrl := strings.Split(corsUrlConn, ",")

		config := cors.Config{
			AllowMethods:     []string{"POST", "GET", "DELETE", "PUT"},
			AllowHeaders:     []string{"Origin", "Content-Type", "X-Requested-With", "Accept"},
			ExposeHeaders:    []string{"Content-Length", "Set-Cookie"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return funk.ContainsString(corsUrl, origin)
			},
			MaxAge: 12 * time.Hour,
		}

		logger.Logger().Info.Printf("Cors urls config is %v", logger.ConvertToString(config))

		r.Use(cors.New(config))
	}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/app")
	})

	r.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))

	r.LoadHTMLFiles("frontend/build/index.html")

	r.GET("/app/*other", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	api.Register(r)

	addresses, err := services.Network().GetAddresses()
	if err != nil {
		return
	}

	var ip4Addresses []string

	for _, address := range addresses {
		if !strings.Contains(address, ":") && address != "127.0.0.1" {
			ip4Addresses = append(ip4Addresses, address)
		}
	}

	port := ":80"

	portVar := os.Getenv("PORT")

	if portVar != "" {
		port = ":" + portVar
	}

	logger.Logger().Info.Printf("Server running at http://%s%s", ip4Addresses[0], port)

	err = r.Run(port)
	if err != nil {
		return
	}

}
