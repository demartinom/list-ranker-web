package global

import "github.com/demartinom/list-ranker-web/pkg/battle"

var Winner battle.Item

func ClearWinner() {
	Winner = battle.Item{}
}
