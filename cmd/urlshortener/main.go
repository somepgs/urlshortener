package main

import (
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	port := ":" + envOrDefault("URLSHORT_PORT", "8080") // Default port is 8080 if not set
	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
