package main

import (
	"fmt"
	"golang/first/utils"

	"github.com/google/uuid"
)

func main() {
	firstUUID := uuid.New()
	fmt.Println(firstUUID)                                // 791e0d38-c0ad-4241-9918-4ac507d92135
	fmt.Println(utils.ReplaceSymbols(firstUUID, "-", "")) // 791e0d38c0ad424199184ac507d92135
	secondUUID := uuid.New()
	fmt.Println(utils.MergeUUID(firstUUID, secondUUID))
	// 791e0d38-c0ad-4241-9918-4ac507d9213523acd18c-4659-45b4-ba8d-c47d895155e5
	fmt.Println(utils.Contains(firstUUID, "e"))  // true
	fmt.Println(utils.Contains(secondUUID, "e")) // true
}
