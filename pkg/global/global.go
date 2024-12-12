package global

var Winner Item
var WinnerPicked = make(chan bool, 1)

type Item struct {
	Name  string
	Score int
}

func (winner *Item) Win() {
	winner.Score++
}

func (loser *Item) Lose(list []*Item, index int) []*Item {
	loser.Score--
	if loser.Score <= -2 {
		list = removeLoser(list, index)
	}
	return list
}

func removeLoser(list []*Item, index int) []*Item {
	return append(list[:index], list[index+1:]...)
}
