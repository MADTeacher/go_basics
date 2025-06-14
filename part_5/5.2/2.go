package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Remove("myDir")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Directory removed")
}
