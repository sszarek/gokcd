package main

import (
	"fmt"
	"os"

	"log"
	"strconv"

	"github.com/sszarek/gokcd/xkcd"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: gokcd <strip_id>")
		os.Exit(0)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("Valid integer should be provided as <strip_id>")
	}

	comic := xkcd.GetComic(id)
	fmt.Println(comic)
}
