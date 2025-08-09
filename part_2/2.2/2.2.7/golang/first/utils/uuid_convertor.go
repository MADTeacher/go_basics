package utils

import (
	"strings"

	"github.com/google/uuid"
)

func ReplaceSymbols(oldUUID uuid.UUID, old, new string) string {
	return strings.ReplaceAll(oldUUID.String(), old, new)
}

func MergeUUID(first, second uuid.UUID) string {
	return first.String() + second.String()
}

func Contains(uuid uuid.UUID, symbol string) bool {
	return strings.ContainsAny(uuid.String(), symbol)
}
