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

type OriginUrlPage struct {
    Origin string
}

func Index(c echo.Context) error {
    return c.Render(200, "index", IndexPage{})
}

func HandleEncode(c echo.Context) error {
	origUrl := c.FormValue("origUrl")

	count, err := database.Count()
	if err != nil {
		return err
	}

	shortKey := encoder.Encode(count)

	err = database.AddUrl(shortKey, origUrl)
	if err != nil {
		log.Fatalf("couldn't save url: %v", err)
	}

	short := "http://localhost:8080/" + shortKey

	fmt.Println(short)

	return c.Render(200, "display", IndexPage{
		ShortUrl: short,
	})
}

func HandleDecode(c echo.Context) error {
    shortKey := c.Param("url")
    
    origin, err := database.GetUrl(shortKey)
    if err != nil {
        log.Fatalf("couldn't get original url: %v", err) 
    }

    fmt.Println(shortKey)

    fmt.Println(origin)

    return c.Render(200, "origin", OriginUrlPage{
        Origin: origin,
    })
}
