package battle

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"slices"

	"github.com/demartinom/list-ranker-web/pkg/global"
	"github.com/gorilla/websocket"
)

func Battle(list []*global.Item, ws *websocket.Conn, results *[]string) {
	for len(list) > 1 {
		battlers, indexes := chooseBattlers(list)
		sendCombatants(battlers, ws)

		<-global.WinnerPicked
		list = battleResult(list, battlers, indexes)
	}
	gameEnding(results, list, ws)
}

func chooseBattlers(list []*global.Item) ([]*global.Item, []int) {
	fighterOneIndex := rand.Intn(len(list))
	fighterTwoIndex := rand.Intn(len(list))

	for fighterOneIndex == fighterTwoIndex {
		fighterTwoIndex = rand.Intn(len(list))
	}

	fighterOne := list[fighterOneIndex]
	fighterTwo := list[fighterTwoIndex]

	combatants := []*global.Item{fighterOne, fighterTwo}
	indexes := []int{fighterOneIndex, fighterTwoIndex}
	return combatants, indexes
}

func battleResult(list []*global.Item, battlers []*global.Item, indexes []int) []*global.Item {
	if global.Winner.Name == battlers[0].Name {
		battlers[0].Win()
		list = battlers[1].Lose(list, indexes[1], &global.Ranking)
	} else {
		list = battlers[0].Lose(list, indexes[0], &global.Ranking)
		battlers[1].Win()
	}

	global.Winner = global.Item{}
	return list
}

func gameEnding(results *[]string, list []*global.Item, ws *websocket.Conn) {
	*results = append(*results, fmt.Sprintf("1. %s", (list)[0].Name))
	slices.Reverse(*results)

	jsonData, err := json.Marshal(ResultsList{"Results", *results})
	if err != nil {
		log.Println("Error marshalling json:", err)
		return
	}

	err = ws.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Println("Error sending message:", err)
	}
}
