package battle

import (
	"encoding/json"

	"github.com/demartinom/list-ranker-web/pkg/global"
)

type PremadeLists struct {
	MessageType string   `json:"messageType"`
	Options     []string `json:"options"`
}

type ReceivedMessage struct {
	MessageType string          `json:"messageType"`
	Data        json.RawMessage `json:"data"`
}

type CombatantsList struct {
	MessageType string         `json:"messageType"`
	Combatants  []*global.Item `json:"combatants"`
}
