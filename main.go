package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"swag-grpc-crud/server"
)

func main() {
	// Start gRPC server
	go func() {
		server.StartGRPCServer()
	}()

	// Start HTTP server
	go func() {
		server.StartHTTPServer()
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down servers...")

	// Clean up tasks if needed
	log.Println("Servers gracefully stopped.")
}
