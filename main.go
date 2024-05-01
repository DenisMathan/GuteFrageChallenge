package main

import (
	"log"

	"github.com/DenisMathan/GuteFrageChallenge/api"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	server := api.NewServer()
	server.Start()
}
