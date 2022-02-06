package main

import (
	"bs-to-scrapper/server/router"
	"bs-to-scrapper/server/services"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/jweslley/localtunnel"
	"github.com/thoas/go-funk"
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

func main() {

	err := godotenv.Load()
	if err != nil {
		return
	}

	r := gin.Default()

	if os.Getenv("GIN_MODE") == "debug" {

		corsUrlConn := os.Getenv("CORS_URL")

		corsUrl := strings.Split(corsUrlConn, ",")

		r.Use(cors.New(cors.Config{
			AllowMethods:     []string{"POST", "GET", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Type", "X-Requested-With", "Accept"},
			ExposeHeaders:    []string{"Content-Length", "Set-Cookie"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return funk.ContainsString(corsUrl, origin)
			},
			MaxAge: 12 * time.Hour,
		}))
	}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/app")
	})

	r.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))

	r.LoadHTMLFiles("frontend/build/index.html")

	r.GET("/app/*other", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.Register(r)

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

	fmt.Printf("Server running at http://%s%s\n", ip4Addresses[0], port)

	err = r.Run(port)
	if err != nil {
		return
	}

}
