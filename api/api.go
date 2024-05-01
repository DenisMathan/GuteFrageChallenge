package api

import (
	"log"
	"net/http"

	"github.com/DenisMathan/codingChallengeGuteFrage/api/router"
	"github.com/DenisMathan/codingChallengeGuteFrage/configurations"
	"github.com/DenisMathan/codingChallengeGuteFrage/database"
)

type server struct {
	app *http.Server
	db  database.SqlHandler
	cfg configurations.Config
}

func NewServer() server {
	cfg := configurations.GetConfig()
	sqlHandler := database.NewSqlHandler(cfg.Database)
	return server{
		app: &http.Server{Addr: ":" + cfg.ServerPort, Handler: router.NewRouter(&sqlHandler)},
		db:  sqlHandler,
		cfg: cfg,
	}
}

func (s server) Start() {
	err := s.app.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal("Http Server stopped unexpected")
		// s.Shutdown()
	} else {
		log.Println("Http Server stopped")
	}
}
