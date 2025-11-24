package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/bobyindra/configs-management-service/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Init Main Rest App
	app, err := config.NewCmsRest()
	if err != nil {
		log.Fatal(err)
	}
	defer app.Redis.Close()
	defer app.Database.Close()

	restServer := config.NewRestServer(app)

	// Interupt signal listener
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Start Rest Server in goroutine
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("Starting Rest server..., listening at %s\n", restServer.ApiAddress())
		if err := restServer.Serve(); err != nil {
			log.Printf("Rest server stopped with error: %v\n", err)
		}
	}()

	// Wait for interupt signal
	<-interrupt
	log.Println("Interupt received. Shutting down server...")

	// Shutdown Rest Server
	if err := restServer.Shutdown(); err != nil {
		log.Printf("Error during REST server shutdown: %v\n", err)
	} else {
		log.Println("Server shutdown completed.")
	}

	// Wait for all goroutines to finish
	wg.Wait()
	log.Println("Server gracefully stopped.")
}
