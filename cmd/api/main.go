package main

import (
	"context"
	"log"
	"net/http"

	"github.com/priince938/app/internal/boot"
	"github.com/priince938/app/internal/database"
	"github.com/priince938/app/internal/provider/config"
	"github.com/priince938/app/internal/router"
)

func main() {

	ctx, cancel := context.WithCancel(boot.NewContext(context.Background()))
	defer cancel()
	// Log that the server has started
	log.Println("Server started on :8080")
	boot.Init()
	database.InitDB(ctx, config.Config.Database)

	// Start the HTTP server on port 8080 and use the router for handling requests
	log.Fatal(http.ListenAndServe(":8080", router.Router(ctx)))
}
