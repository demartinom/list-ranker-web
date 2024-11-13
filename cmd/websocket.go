package main

import (
	"log"
	"net/http"

	"github.com/demartinom/list-ranker-web/pkg/battle"
	"github.com/gorilla/websocket"
)

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

	battle.SendBattleOptions(ws)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Error reading message %v\n", err)
		}
		log.Printf("Received message : %s\n", msg)

		// if err = ws.WriteMessage(websocket.TextMessage, msg); err != nil {
		// 	log.Printf("Error: %v\n", err)
		// 	break
		// }
	}
}
