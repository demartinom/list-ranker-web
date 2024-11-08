package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
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

// Create instance of Upgrader struct to upgrade http connection to websocket
var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	log.Println("Connected Successfully")

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading message %v\n", err)
		}
		log.Printf("Received message : %s\n", msg)

		if err = ws.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Printf("Error: %v\n", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	fmt.Printf("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting sever %v\n", err)
	}
}
