package battle

import "fmt"

func Battle(list []Item) {
func Battle(list []Item, ws *websocket.Conn) {
	battlers, _ := chooseBattlers(list)

	for _, v := range battlers {
		fmt.Println(v)
	}
}
