package battle

import "encoding/json"

type Item struct {
	Name  string
	Score int
}

type PremadeLists struct {
	MessageType string   `json:"messageType"`
	Options     []string `json:"options"`
}

type ReceivedMessage struct {
	MessageType string          `json:"messageType"`
	Data        json.RawMessage `json:"data"`
}
