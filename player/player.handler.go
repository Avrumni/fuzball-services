package player

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	players := GetAll()
	json.NewEncoder(w).Encode(players)
}

func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	player, _ := GetById(params["id"])
	json.NewEncoder(w).Encode(player)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var player Player
	_ = json.NewDecoder(r.Body).Decode(&player)

	player = Create(player)

	//json.NewEncoder(w).Encode(player)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	DeleteById(params["id"])
}