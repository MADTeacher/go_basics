package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Mkdir("myDir", 0777)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Directory created")
}
