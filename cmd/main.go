package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func serverStart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server started")
}

type Message struct {
	Text string `json:"text"`
}

// Set headers for all handleFuncs to enable CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight (OPTIONS) request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func test(w http.ResponseWriter, r *http.Request) {
	data := Message{Text: "Hello from the backend"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	mux := http.NewServeMux()

	fmt.Println("Starting server on port 8080")

	corsHandler := enableCORS(mux)

	mux.HandleFunc("/", serverStart)
	mux.HandleFunc("/message", test)
	http.ListenAndServe(":8080", corsHandler)
}
