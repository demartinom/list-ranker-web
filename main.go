package main

import (
	"fmt"
	"net/http"
)

func serverStart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server started")
}

func main() {
	fmt.Println("Starting server on port 8080")

	http.HandleFunc("/", serverStart)
	http.ListenAndServe(":8080", nil)
}
