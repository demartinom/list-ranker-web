package battle

import (
	"math/rand"

	"github.com/demartinom/list-ranker-web/pkg/global"
	"github.com/gorilla/websocket"
)

func Battle(list []global.Item, ws *websocket.Conn) {
	for len(list) > 1 {
		battlers, _ := chooseBattlers(list)
		sendCombatants(battlers, ws)
		<-global.WinnerPicked
		battleResult(battlers)
	}
}

func chooseBattlers(list []global.Item) ([]global.Item, []int) {
	fighterOneIndex := rand.Intn(len(list))
	fighterTwoIndex := rand.Intn(len(list))

	for fighterOneIndex == fighterTwoIndex {
		fighterTwoIndex = rand.Intn(len(list))
	}

	fighterOne := list[fighterOneIndex]
	fighterTwo := list[fighterTwoIndex]

	combatants := []global.Item{fighterOne, fighterTwo}
	indexes := []int{fighterOneIndex, fighterTwoIndex}
	return combatants, indexes
}

func battleResult(battlers []global.Item) {
	if global.Winner.Name == battlers[0].Name {
		battlers[0].Win()
		battlers[1].Lose()
	} else {
		battlers[0].Lose()
		battlers[1].Win()
	}
	global.Winner = global.Item{}
}
