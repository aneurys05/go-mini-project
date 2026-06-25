package main

import (
	"log"
	"net/http"
	"petpal/internal/db"
	"petpal/internal/router"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	r := router.New(database)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
