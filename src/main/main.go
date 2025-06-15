package main

import application "github.com/brianmorais/go-book-store-api/src/application"

func main() {
	server := application.NewWebServer()
	server.Serve()
}
