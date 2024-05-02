package database

import (
	"log"
	"testing"

	"github.com/DenisMathan/GuteFrageChallenge/configurations"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestBuildDsn(t *testing.T) {
	cfg := configurations.Config{
		Database: configurations.Database{
			Host:     "localhost",
			Port:     3307,
			User:     "admin",
			Password: "root",
			Name:     "guteFrage",
		},
	}
	expected := "admin:root@tcp(localhost:3307)/guteFrage?charset=utf8mb4&parseTime=True&loc=Local"
	result := builddsn(cfg.Database)
	if result != expected {
		t.Errorf("builddsn returned %s, expected %s", result, expected)
	}
}

func TestGetTodos(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	//onlyDone
	mock.ExpectQuery("SELECT \\* FROM todos WHERE done = true LIMIT \\? OFFSET \\?").WithArgs(10, 0).WillReturnRows(sqlmock.NewRows([]string{"id", "description", "done"}).AddRow(0, "John Doe", true))
	//all
	mock.ExpectQuery("SELECT \\* FROM todos LIMIT \\? OFFSET \\?").WithArgs(10, 0).WillReturnRows(sqlmock.NewRows([]string{"id", "description", "done"}).AddRow(0, "John Doe", true).AddRow(1, "John Doe", false))

	mockHandler := SqlHandler{db: db}

	// Testfall: onlyDone
	mockHandler.GetTodos(true, 10, 0)
	// log.Println(todos)

	// Testfall: Alle Todos abrufen
	mockHandler.GetTodos(false, 10, 0)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
