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

type CombatantsList struct {
	MessageType string `json:"messageType"`
	Combatants  []Item `json:"combatants"`
}

type BattleResult struct {
	MessageType string `json:"messageType"`
	Winner      Item   `json:"winner"`
}
