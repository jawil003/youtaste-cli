package scrapper

import "github.com/go-rod/rod"

func ScreenshotOnError(page *rod.Page) {
	page.MustScreenshot("error.png")
}
