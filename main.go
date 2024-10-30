package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize Kafka 
	InitKafkaWriter("localhost:9092", "default-topic")

	// Router setup
	router := NewRouter()

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
