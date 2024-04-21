package handler

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"github.com/Roddyck/go_url_short/encoder"
	"github.com/Roddyck/go_url_short/pkg/database"
)

type IndexPage struct {
	ShortUrl string
}

func HandleEncode(c echo.Context) error {
	origUrl := c.FormValue("origUrl")

	shortKey := encoder.Encode()

	err := database.AddUrl(shortKey, origUrl)
	if err != nil {
		log.Fatalf("couldn't save url: %v", err)
	}

	short := "http://localhost:8080/" + shortKey

    fmt.Println(short)

	return c.Render(200, "display", IndexPage{
		ShortUrl: short,
	})
}

func Index(c echo.Context) error {
	return c.Render(200, "index", IndexPage{})
}
