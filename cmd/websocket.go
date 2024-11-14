package main

import (
	"encoding/json"
	"fmt"
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
		var message battle.ReceivedMessage

		_, msg, err := ws.ReadMessage()
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("Error unmarshalling:", err)
		}
		if err != nil {
			log.Printf("Error reading message %v\n", err)
		}

		switch message.MessageType {
		case "Premade List":
			var listChoice battle.PremadeList
			if err := json.Unmarshal(message.Data, &listChoice); err != nil {
				log.Println("Error unmarshalling:", err)
			}
			fmt.Println(listChoice.List)
		case "Custom List":
			var customList battle.CustomList
			if err := json.Unmarshal(message.Data, &customList); err != nil {
				log.Println("Error unmarshalling:", err)
			}
			fmt.Println(customList.List)
		}

		// if err = ws.WriteMessage(websocket.TextMessage, msg); err != nil {
		// 	log.Printf("Error: %v\n", err)
		// 	break
		// }
	}
}
