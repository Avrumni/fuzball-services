package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/golang-migrate/migrate/source/file"
	"motome.com.au/fuzball-services/player"
	"motome.com.au/fuzball-services/db"
)

// our main function
func main() {
	db.Connect()

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