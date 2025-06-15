package main

import (
	application "github.com/brianmorais/go-book-store-api/src/application"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	server := application.NewWebServer()
	server.Serve()
}
