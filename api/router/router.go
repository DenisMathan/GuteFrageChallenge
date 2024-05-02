package router

import (
	"context"
	"io"
	"net/http"

	"github.com/DenisMathan/GuteFrageChallenge/entities"
	"github.com/gorilla/mux"
)

type SqlHandler interface {
	GetTodos(done bool, pagination int, nth int) []entities.Todo
}

func NewRouter(sqlHandler SqlHandler) *mux.Router {
	router := mux.NewRouter()
	router.Use(CORS)
	router.Use(func(next http.Handler) http.Handler { return fillContext(next, sqlHandler) })
	requests(router)
	return router
}

func requests(router *mux.Router) {
	router.HandleFunc("/", routeBase)
	router.HandleFunc("/todos", getTodos)
}

func routeBase(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "router Works")
}

// Filling Context to have access to certain objects during the requestworkflow
func fillContext(next http.Handler, sqlHandler SqlHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values := map[string]interface{}{
			"sqlHandler": sqlHandler,
		}
		newCtx := context.WithValue(r.Context(), "values", values)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
