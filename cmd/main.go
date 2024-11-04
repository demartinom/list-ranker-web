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

func test(w http.ResponseWriter, r *http.Request) {
	data := Message{Text: "Hello from the backend"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	fmt.Println("Starting server on port 8080")

	http.HandleFunc("/", serverStart)
	http.HandleFunc("/message", test)
	http.ListenAndServe(":8080", nil)
}
