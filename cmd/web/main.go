package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/erkindilekci/debt-dash/pkg/db"
	"github.com/erkindilekci/debt-dash/pkg/routes"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Listening and serving the project on port number %s\n", port)
	err = http.ListenAndServe(":"+port, routes.Routes())
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
