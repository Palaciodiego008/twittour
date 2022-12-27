package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers is the function that handles the routes
func Handlers() {
	router := mux.NewRouter()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3001"
	}

	corsPermit := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, corsPermit))

}
