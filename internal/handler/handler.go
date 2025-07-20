package handler

import (
	"log"
	"net/http"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	return mux
}
