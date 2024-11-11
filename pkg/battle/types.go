package battle

type Item struct {
	Name  string
	Score int
}

type ListChoices struct {
	MessageType string   `json:"messageType"`
	Options     []string `json:"options"`
}
