package handler

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Roddyck/go_url_short/encoder"
	"github.com/Roddyck/go_url_short/pkg/database"
)

type URLRequest struct {
	URL string `json:"url"`
}

type URLStorage struct {
	urls map[string]string
}

var us URLStorage = URLStorage{urls: make(map[string]string)}

var UseDB = flag.Bool("d", false, "use either postgers or memory")

func init() {
	flag.Parse()
}

func HandleEncode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST / request")
	var url URLRequest

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		log.Fatalf("could not decode url: %v", err)
	}

	shortKey := encoder.Encode()

	if *UseDB {
		err := database.AddUrl(shortKey, url.URL)
		if err != nil {
			log.Fatalf("could not save url to db: %v", err)
		}
	} else {
		us.urls[shortKey] = url.URL
	}

	shortUrl := "http://localhost:8080/" + shortKey

	json.NewEncoder(w).Encode(shortUrl)
}

func HandleDecode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /{url} request")

	shortKey := r.PathValue("url")

	var origin string

	if *UseDB {
		origUrl, err := database.GetUrl(shortKey)
		if err != nil {
            log.Fatalf("error retriving url from db: %v", err)
		}
        origin = origUrl
	} else {
		origUrl, found := us.urls[shortKey]
		if !found {
			log.Fatalln("shortend key not found")
		}
        origin = origUrl
	}

	json.NewEncoder(w).Encode(origin)
}
