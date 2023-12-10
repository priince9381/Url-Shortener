package main

import (
	"log"
	"net/http"

	"github.com/priince938/app/internal/router"
)

func main() {
	// Log that the server has started
	log.Println("Server started on :8080")

	// Start the HTTP server on port 8080 and use the router for handling requests
	log.Fatal(http.ListenAndServe(":8080", router.Router()))
}
