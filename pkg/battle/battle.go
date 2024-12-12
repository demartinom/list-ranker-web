package battle

import (
	"fmt"
	"math/rand"

	"github.com/demartinom/list-ranker-web/pkg/global"
	"github.com/gorilla/websocket"
)

func Battle(list []*global.Item, ws *websocket.Conn) {
	for len(list) > 1 {
		battlers, indexes := chooseBattlers(list)
		sendCombatants(battlers, ws)

		<-global.WinnerPicked
		list = battleResult(list, battlers, indexes)

		fmt.Println("")
		for _, item := range list {
			fmt.Println(item)
		}
	}
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
