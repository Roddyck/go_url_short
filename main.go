package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Roddyck/go_url_short/handler"
	"github.com/Roddyck/go_url_short/pkg/database"
)

func main() {
    flag.Parse()

    name := "urls"

    if *handler.UseDB {
        err := database.InitDB(name)
        if err != nil {
            log.Fatalf("could not initialize db: %v", err)
        }
        fmt.Println("db running")
    }

    mux := http.NewServeMux()
    mux.HandleFunc("POST /", handler.HandleEncode)
    mux.HandleFunc("GET /{url}", handler.HandleDecode)

    fmt.Println("listening on port :8080")
    http.ListenAndServe(":8080", mux)
}


