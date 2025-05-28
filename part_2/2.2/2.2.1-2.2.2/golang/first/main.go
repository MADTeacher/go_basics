package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	myUUID := uuid.New()
	fmt.Println(myUUID)
}
