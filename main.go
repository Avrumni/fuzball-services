package main

import (
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"fmt"
	"motome.com.au/fuzball-services/player"
)

// our main function
func main() {
	dbUrl := os.Getenv("DATABASE_URL")
	fmt.Println(dbUrl)
	db, err := sql.Open("postgres", dbUrl)
	if (err != nil) {
		print(1, err.Error())
		return
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if (err != nil) {
		print(2, err.Error())
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if (err != nil) {
		print(3, err.Error())
		return
	}

	m.Up()

	router := mux.NewRouter()
	router.HandleFunc("/player", player.GetAllHandler).Methods("GET")
	router.HandleFunc("/player/{id}", player.GetByIdHandler).Methods("GET")
	router.HandleFunc("/player", player.CreateHandler).Methods("POST")
	router.HandleFunc("/player/{id}", player.DeleteHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

type Team struct {
	Player1 player.Player `json:"player1,omitempty"`
	Player2 player.Player `json:"player2,omitempty"`
}

type Match struct {
	TeamA Team `json:"teamA,omitempty"`
	TeamB Team `json:"teamB,omitempty"`
}