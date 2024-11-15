package battle

import (
	"fmt"
	"math/rand"

	"github.com/gorilla/websocket"
)

func Battle(list []Item, ws *websocket.Conn) {
	battlers, _ := chooseBattlers(list)
	sendCombatants(battlers, ws)
	for _, v := range battlers {
		fmt.Println(v)
	}
}

func chooseBattlers(list []Item) ([]Item, []int) {
	fighterOneIndex := rand.Intn(len(list))
	fighterTwoIndex := rand.Intn(len(list))

	for fighterOneIndex == fighterTwoIndex {
		fighterTwoIndex = rand.Intn(len(list))
	}

	fighterOne := list[fighterOneIndex]
	fighterTwo := list[fighterTwoIndex]

	combatants := []Item{fighterOne, fighterTwo}
	indexes := []int{fighterOneIndex, fighterTwoIndex}
	return combatants, indexes
}
