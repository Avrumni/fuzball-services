package main

import (
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"fmt"
)

var players []Player

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

	players = append(players, Player{ID: "1", Firstname: "Elliot", Lastname: "Smith"})

	router := mux.NewRouter()
	router.HandleFunc("/player", GetPlayers).Methods("GET")
	router.HandleFunc("/player/{id}", GetPlayer).Methods("GET")
	router.HandleFunc("/player", CreatePlayer).Methods("POST")
	router.HandleFunc("/player/{id}", DeletePlayer).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

type Player struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
}

type Team struct {
	Player1 Player `json:"player1,omitempty"`
	Player2 Player `json:"player2,omitempty"`
}

type Match struct {
	TeamA Team `json:"teamA,omitempty"`
	TeamB Team `json:"teamB,omitempty"`
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range players {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var player Player
	_ = json.NewDecoder(r.Body).Decode(&player)
	player.ID = params["id"]
	players = append(players, player)
	json.NewEncoder(w).Encode(players)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range players {
		if item.ID == params["id"] {
			players = append(players[:index], players[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(players)
}