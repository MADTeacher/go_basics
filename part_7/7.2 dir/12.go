package main

import (
	"fmt"
	"os"
)

func main() {

	pathToDir := "C:\\Go"

	if _, err := os.Stat(pathToDir); os.IsNotExist(err) {
		// обработка ситуации, когда директории не существует
		fmt.Println("Directory isn't exist")
	} else {
		fmt.Println("Directory already exists")
	}
}
