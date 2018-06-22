package team

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	teams := GetAll()
	json.NewEncoder(w).Encode(teams)
}

func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	team, _ := GetById(params["id"])
	json.NewEncoder(w).Encode(team)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var team Team
	_ = json.NewDecoder(r.Body).Decode(&team)

	updatedTeam := Create(team)

	json.NewEncoder(w).Encode(updatedTeam)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	DeleteById(params["id"])
}