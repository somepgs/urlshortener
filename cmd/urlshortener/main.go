package main

import (
	"github.com/somepgs/urlshortener/internal/handler"
	"log"
	"net/http"
	"os"
)

// envOrDefault retrieves the value of an environment variable or returns a default value if the variable is not set.
func envOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// main initializes the HTTP server and handles requests.
func main() {
	port := ":" + envOrDefault("URLSHORT_PORT", "8080")
	log.Printf("Starting server on port %s", port)

	router := handler.SetupRoutes()

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
