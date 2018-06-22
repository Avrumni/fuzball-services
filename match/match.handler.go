package match

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	matchs := GetAll()
	json.NewEncoder(w).Encode(matchs)
}

func GetByIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	match, _ := GetById(params["id"])
	json.NewEncoder(w).Encode(match)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var match Match
	_ = json.NewDecoder(r.Body).Decode(&match)

	updatedMatch := Create(match)

	json.NewEncoder(w).Encode(updatedMatch)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	DeleteById(params["id"])
}