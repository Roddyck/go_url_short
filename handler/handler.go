package handler

import (
	"github.com/labstack/echo/v4"
)


type IndexPage struct {
    shortUrl string
}

func Index(c echo.Context) error {
    return c.Render(200, "index", IndexPage{
        shortUrl: "",
    })
}
