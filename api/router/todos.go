package router

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func getTodos(w http.ResponseWriter, request *http.Request) {
	sqlH := request.Context().Value("values").(map[string]interface{})["sqlHandler"].(SqlHandler)
	query := request.URL.Query()
	onlyDone := query.Get("onlyDone") == "true"
	pagination, err := strconv.Atoi(query.Get("pagination"))
	if err != nil {
		log.Println(err)
		http.Error(w, "pagination is missing", http.StatusBadRequest)
		return
	}
	nth, err := strconv.Atoi(query.Get("nth"))
	if err != nil {
		log.Println(err)
		http.Error(w, "nth (paginationstep) is missing)", http.StatusBadRequest)
		return
	}
	todos := sqlH.GetTodos(onlyDone, pagination, nth)
	json.NewEncoder(w).Encode(todos)
}
