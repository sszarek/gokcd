package xkcd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

// GetComic returns single comic
func GetComic(id int) *Comic {
	name := fmt.Sprintf("xkcd-%d.json", id)
	if _, err := os.Stat(name); os.IsNotExist(err) {
		indexComic(id, name)
	}

	return loadFromIndex(name)
}

func indexComic(id int, name string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("Error creating index file: %s", err)
	}

	resp, err := http.Get(ApiURL)
	if err != nil {
		resp.Body.Close()
		log.Fatalf("Error fetching the comic: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		log.Fatalf("Server returned non OK status code: %d", resp.StatusCode)
	}

	io.Copy(file, resp.Body)
	resp.Body.Close()
	file.Close()
}

func loadFromIndex(name string) *Comic {
	file, err := os.Open(name)
	if err != nil {
		file.Close()
		log.Fatalf("Error reading index file: %s", err)
	}

	var comic Comic
	err = json.NewDecoder(file).Decode(&comic)
	if err != nil {
		file.Close()
		log.Fatalf("Error unmarshalling comic: %s", err)
	}

	file.Close()
	return &comic
}
