package xkcd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Comic strip representation
type Comic struct {
	Month      string
	ComicID    int32 `json:"num"`
	Link       string
	News       string
	Title      string
	Transcript string
	Alt        string
	ImgURL     string `json:"img"`
	Day        string
}

// ApiURL is URL of XKCD api
const ApiURL = "https://xkcd.com/info.0.json"

// GetAll gets XKCD comics
func GetAll() {
	resp, err := http.Get(ApiURL)
	if err != nil {
		resp.Body.Close()
		log.Fatalf("Error fetching the comic: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("Server returned non OK status code: %d", resp.StatusCode)
	}

	var comic Comic
	err = json.NewDecoder(resp.Body).Decode(&comic)
	if err != nil {
		resp.Body.Close()
		log.Fatalf("Error unmarshalling the comic: %s", err)
	}

	fmt.Printf("Title: %s\n", comic.Title)
}
