package main

import (
	"fmt"
	"log"
	"net/http"
)

// // Set headers for all handleFuncs to enable CORS
// func enableCORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 		// Handle preflight (OPTIONS) request
// 		if r.Method == http.MethodOptions {
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

func main() {
	http.HandleFunc("/ws", handleConnections)

	fmt.Printf("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting sever %v\n", err)
	}
}
