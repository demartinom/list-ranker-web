package battle

import "fmt"

func Battle(list []Item) {
	battlers, _ := chooseBattlers(list)

	for _, v := range battlers {
		fmt.Println(v)
	}
}
