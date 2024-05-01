package main

import (
	"fmt"
	"log"
	"html/template"
    "io"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

	"github.com/Roddyck/go_url_short/handler"
	"github.com/Roddyck/go_url_short/pkg/database"
)

type Templates struct {
    templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
    return &Templates{
        templates: template.Must(template.ParseGlob("views/*.html")),
    }
}

func main() {
    name := "urls"

    err := database.InitDB(name)
    if err != nil {
        log.Fatalf("could not initialize db: %v", err)
    }
    fmt.Println("db running")

    e := echo.New()
    e.Use(middleware.Logger())

    e.Renderer = NewTemplate()

    e.Static("/css", "css")

    e.GET("/", handler.Index)
    e.POST("/encode", handler.HandleEncode)
    e.GET("/:url", handler.HandleDecode)

    e.Logger.Fatal(e.Start(":8080"))
}


