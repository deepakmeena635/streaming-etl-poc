package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/myname/myrepo/internal/api"
	"github.com/myname/myrepo/pkg/server"
	"github.com/myname/myrepo/pkg/settings"
)

// TODO: Define the period for updating parameters in DynamoDB
const dynamoDBUpdatePeriod = 1 * time.Minute

func main() {
	// TODO: Load server settings from configuration file
	settings, err := settings.Load("config/dev.yaml")
	if err != nil {
		log.Fatalf("Error loading settings: %s", err)
	}

	// TODO: Initialize the API with a DynamoDB client and a cooling controller
	api := api.New(settings)

	// TODO: Initialize the server with the API and settings
	server := server.New(api, settings)

	// TODO: Start a background goroutine to periodically update parameters in DynamoDB
	go updateDynamoDBPeriodically(settings)

	// TODO: Start the server
	go func() {
		log.Printf("Starting server on %s\n", settings.ServerAddr)
		err = server.ListenAndServe()
		if err != nil {
			log.Fatalf("Error starting server: %s", err)
		}
	}()

	// Wait for SIGINT or SIGTERM to gracefully shutdown the server
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown

	log.Println("Shutting down server...")
	server.Shutdown()
	log.Println("Server stopped.")
}

func updateDynamoDBPeriodically(settings *settings.Settings) {
	for {
		// TODO: Implement the logic to update parameters in DynamoDB
		log.Println("Updating parameters in DynamoDB...")
		time.Sleep(dynamoDBUpdatePeriod)
	}
}
