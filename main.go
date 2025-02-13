package main

import (
	"log"
	"net/http"

	"Github.com/Ekaloi/library/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ser := server.NewServer()
	http.HandleFunc("GET /title", ser.SearchBookHandler)
	http.HandleFunc("GET /volumeid/", ser.SearchBookByIDHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Error starting server:", err)
    }

}
