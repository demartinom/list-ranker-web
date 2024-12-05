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

func (loser *Item) Lose() {
	loser.Score--
}
