package main

import (
	"log"
	"twittour/db"
	"twittour/handlers"
)

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("No connection to the database")
		return
	}

	handlers.Handlers()

}
