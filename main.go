package main

import (
	"log"

	"github.com/DenisMathan/codingChallengeGuteFrage/api"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	server := api.NewServer()
	server.Start()
}
