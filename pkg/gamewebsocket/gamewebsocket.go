package gamewebsocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/demartinom/list-ranker-web/pkg/battle"
	"github.com/demartinom/list-ranker-web/pkg/global"
	"github.com/gorilla/websocket"
)

// Create instance of Upgrader struct to upgrade http connection to websocket
var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
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
			log.Printf("Error reading message: %v\n", err)
			break
		}

		if len(msg) == 0 {
			log.Println("Received empty message, skipping...")
		}

		var message battle.ReceivedMessage
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Printf("Error unmarshalling message: %v\n", err)
			continue
		}

		switch message.MessageType {
		case "Premade List":
			var listChoice string
			if err := json.Unmarshal(message.Data, &listChoice); err != nil {
				log.Println("Error unmarshalling:", err)
			}
			battleList := battle.ReadCSV(listChoice)
			go battle.Battle(battleList, ws, &global.Ranking)

		case "Custom List":
			var customList []string
			if err := json.Unmarshal(message.Data, &customList); err != nil {
				log.Println("Error unmarshalling:", err)
			}
			battleList := battle.ReadCustom(customList)
			go battle.Battle(battleList, ws, &global.Ranking)
		case "Result":
			if err := json.Unmarshal(message.Data, &global.Winner); err != nil {
				log.Println("Error unmarshalling:", err)
			}
			global.WinnerPicked <- true
		}
	}
}
