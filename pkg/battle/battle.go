package battle

import (
	"fmt"
	"math/rand"

	"github.com/demartinom/list-ranker-web/pkg/global"
	"github.com/gorilla/websocket"
)

func Battle(list []global.Item, ws *websocket.Conn) {
	battlers, _ := chooseBattlers(list)
	sendCombatants(battlers, ws)
	<-global.WinnerPicked
	fmt.Println(global.Winner)
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
