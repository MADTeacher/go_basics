package network

import "encoding/json"

type Command string

type Message struct {
	Cmd     Command         `json:"command"`
	Payload json.RawMessage `json:"payload,omitempty"`
}
