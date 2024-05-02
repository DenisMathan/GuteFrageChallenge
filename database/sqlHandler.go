package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/DenisMathan/GuteFrageChallenge/configurations"
	"github.com/DenisMathan/GuteFrageChallenge/entities"
)

type SqlHandler struct {
	db *sql.DB
}

func builddsn(cfg configurations.Database) string {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)
	log.Println(dsn)
	return dsn
}

func NewSqlHandler(cfg configurations.Database) SqlHandler {
	dsn := builddsn(cfg)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return SqlHandler{
		db: db,
	}
}

// pagination -> amount elements per request
// nth -> position in elementlist from which to start sending
func (handler *SqlHandler) GetTodos(onlyDone bool, pagination int, nth int) []entities.Todo {
	var rows *sql.Rows
	var err error
	if onlyDone {
		rows, err = handler.db.Query("SELECT * FROM todos WHERE done = true LIMIT ? OFFSET ?", pagination, pagination*nth)
		if err != nil {
			log.Println(err)
		}
	} else {
		rows, err = handler.db.Query("SELECT * FROM todos LIMIT ? OFFSET ?", pagination, pagination*nth)
		if err != nil {
			log.Println(err)
		}
	}
	if err != nil {
		log.Println(err)
	}
	results := []entities.Todo{}
	if rows != nil {
		for rows.Next() {
			var id uint
			var description string
			var done bool
			err = rows.Scan(&id, &description, &done)
			if err != nil {
				panic(err)
			}
			results = append(results, entities.Todo{ID: id, Description: description, Done: done})
		}

		defer rows.Close()
	}
	log.Println(results)
	return results
}

// TODO CREATE
// TODO UPDATE
// TODO DELETE
