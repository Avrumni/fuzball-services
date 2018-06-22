package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/golang-migrate/migrate/source/file"
	"motome.com.au/fuzball-services/player"
	"motome.com.au/fuzball-services/db"
	"motome.com.au/fuzball-services/team"
	"motome.com.au/fuzball-services/match"
)

// our main function
func main() {
	db.Connect()

	router := mux.NewRouter()
	router.HandleFunc("/player", player.GetAllHandler).Methods("GET")
	router.HandleFunc("/player/{id}", player.GetByIdHandler).Methods("GET")
	router.HandleFunc("/player", player.CreateHandler).Methods("POST")
	router.HandleFunc("/player/{id}", player.DeleteHandler).Methods("DELETE")

	router.HandleFunc("/match", match.GetAllHandler).Methods("GET")
	router.HandleFunc("/match/{id}", match.GetByIdHandler).Methods("GET")
	router.HandleFunc("/match", match.CreateHandler).Methods("POST")
	router.HandleFunc("/match/{id}", match.DeleteHandler).Methods("DELETE")

	router.HandleFunc("/team", team.GetAllHandler).Methods("GET")
	router.HandleFunc("/team/{id}", team.GetByIdHandler).Methods("GET")
	router.HandleFunc("/team", team.CreateHandler).Methods("POST")
	router.HandleFunc("/team/{id}", team.DeleteHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

//type Team struct {
//	Player1 player.Player `json:"player1,omitempty"`
//	Player2 player.Player `json:"player2,omitempty"`
//}
//
//type Match struct {
//	TeamA Team `json:"teamA,omitempty"`
//	TeamB Team `json:"teamB,omitempty"`
//}