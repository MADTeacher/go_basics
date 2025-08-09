package main

import (
	f "fmt"
	pack "golang/first/another_package"
	u "golang/first/utils"

	"golang/second"

	id "github.com/google/uuid"
)

func main() {
	firstUUID := id.New()
	f.Println(u.ReplaceSymbols(firstUUID, "-", "")) // 8912acef6d3d40aa94f21333faa1c679
	f.Println(pack.Add(4, 7))                       // 11
	f.Println(second.AddInt(3, 5))                  // 8
	f.Println(second.DivFloat(54.48, 3.66))         // 14.885245901639342
}
