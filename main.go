package main

import (
	"log"

	"github.com/DenisMathan/GuteFrageChallenge/api"
	"github.com/DenisMathan/GuteFrageChallenge/configurations"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg, err := GetConfig()
	if err != nil {
		panic(err)
	}
	server := api.NewServer(cfg)
	server.Start()
}

func GetConfig() (configurations.Config, error) {
	err := godotenv.Load("./.env")

	cnf := configurations.Config{}
	envconfig.Process("", &cnf)
	return cnf, err
}
