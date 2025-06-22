package network

import "encoding/json"

// Объявляем тип команды
type Command string

// Базовая структура сообщения
// Cmd - тип команды
// Payload - полезная нагрузка, которая должна быть
// сериализована в JSON ([]byte). Может быть nil
type Message struct {
	Cmd     Command         `json:"command"`
	Payload json.RawMessage `json:"payload,omitempty"`
}
