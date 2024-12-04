package global

var Winner Item

func ClearWinner() {
	Winner = Item{}
}

type Item struct {
	Name  string
	Score int
}
