package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func OkHandler(w http.ResponseWriter, r *http.Request) {
	payload := map[string]interface{}{
		"status":  "OK",
		"service": "chats",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}

	router := mux.NewRouter()

	router.Use(loggingMiddleware)

	router.HandleFunc("/health-check", OkHandler).Methods("GET")
	router.HandleFunc("/ready-check", OkHandler).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Lets go!")
	log.Fatal(srv.ListenAndServe())
}
