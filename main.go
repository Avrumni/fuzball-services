package main

import (
	"log"
	"net/http"
	_ "github.com/lib/pq"
	_ "github.com/golang-migrate/migrate/source/file"
	"motome.com.au/fuzball-services/db"
	"os"
	"github.com/gorilla/mux"
	"motome.com.au/fuzball-services/player"
	"motome.com.au/fuzball-services/match"
	"motome.com.au/fuzball-services/team"
)

// our main function
func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	println("Starting on port:" + port)
	db.Connect()

	router := mux.NewRouter()
	router.HandleFunc("/player", player.GetAllHandler).Methods("GET")
	router.HandleFunc("/player/{id}", player.GetByIdHandler).Methods("GET")
	//TODO Rethink API naming
	router.HandleFunc("/player/name/{name}", player.GetByNameHandler).Methods("GET")
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

	router.Use(loggingMiddleware)

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	log.Fatal(http.ListenAndServe(":" + port, router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}