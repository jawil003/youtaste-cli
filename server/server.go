package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
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

	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	err := r.Run()
	if err != nil {
		return
	}

}

func run(c *cli.Context) {

	shouldAuth := c.Bool("auth")

	var order OrderWithRes

	jsonFile, err := os.Open("order.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &order)
	if err != nil {
		return
	}

	u := launcher.NewUserMode().
		MustLaunch()

	controller := rod.New().ControlURL(u)

	page := controller.MustConnect().NoDefaultDevice().MustPage("https://youtaste.com/")

	if shouldAuth {

		err = godotenv.Load()
		if err != nil {
			return
		}

		err = page.MustElementR("#navigation a", "Einloggen").Click(proto.InputMouseButtonLeft)
		if err != nil {
			return
		}

		page.MustWaitNavigation()

		phoneInput := page.MustElement("input[placeholder=\"Telefonnummer\"]")
		phoneNumber := os.Getenv("PHONE")
		phoneInput.MustInput(phoneNumber)

		passwordInput := page.MustElement("input[placeholder=\"Passwort\"]")
		passwordInput.MustInput(os.Getenv("PASSWORD"))

		err = page.MustElementR("button", "Anmelden").Click(proto.InputMouseButtonLeft)
		if err != nil {
			return
		}

		page.MustWaitNavigation()
	}

	searchInput := page.MustElement("input#search-restaurant-input")

	searchInput.MustInput(order.Restaurant)
	searchInput.MustPress(input.Enter)

	err = page.MustElement("#restaurantList a").Click(proto.InputMouseButtonLeft)
	if err != nil {
		return
	}

	page.MustWaitNavigation()

	for _, product := range order.Order {
		err = searchForProduct(product, page)
		if err != nil {
			return
		}
	}

	page.MustWaitNavigation()
}

func searchForProduct(p Product, page *rod.Page) error {
	element, err := page.ElementR("#search-content-div a", p.Name)
	if err != nil {
		return err
	}

	err = element.Click(proto.InputMouseButtonLeft)

	for _, variant := range p.Variant {
		regex := fmt.Sprintf("/\\s*%s\\s*/gmi", variant)
		element = page.MustElementR("#productModalForm div.text-black", regex)
		err := element.Click(proto.InputMouseButtonLeft)
		if err != nil {
			return err
		}
	}

	page.MustElement("input[type=\"submit\"]").MustClick()

	if err != nil {
		return err
	}

	return nil
}
