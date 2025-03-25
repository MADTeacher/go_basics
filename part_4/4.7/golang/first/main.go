package main

import (
	f "fmt"
	pack "golang/first/another_package"
	u "golang/first/utils"

	id "github.com/google/uuid"
)

func main() {
	firstUUID := id.New()
	f.Println(u.ReplaceSymbols(firstUUID, "-", "")) // 45f09a89584b4f66936cf3368e3f634b
	print(pack.Add(4, 7))                           // 11
}
