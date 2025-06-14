package main

import (
	"fmt"
)

var VersionType = "dev"
var Version = "0.0.1"

func main() {
	fmt.Printf("Version: %s (%s)\n", Version, VersionType)
}
