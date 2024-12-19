package global

import "fmt"

var Winner Item
var WinnerPicked = make(chan bool, 1)
var Ranking []string

type Item struct {
	Name  string
	Score int
}

func (winner *Item) Win() {
	winner.Score++
}

func (loser *Item) Lose(list []*Item, index int, results *[]string) []*Item {
	loser.Score--
	if len(list) == 2 {
		list = removeLoser(list, index, results)
	}
	if loser.Score <= -2 {
		list = removeLoser(list, index, results)
	}
	return list
}

func removeLoser(list []*Item, index int, results *[]string) []*Item {
	placement := fmt.Sprintf("%d: %s", len(list), (list)[index].Name)
	*results = append(*results, placement)
	return append(list[:index], list[index+1:]...)
}
