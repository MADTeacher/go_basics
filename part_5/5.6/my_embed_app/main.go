package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
)

//go:embed embed_config.json
var configData []byte

type Config struct {
	Debug   bool   `json:"debug"`
	Offset  int    `json:"offset"`
	AppName string `json:"app_name"`
	Version string `json:"version"`
}

func getConfig() Config {
	var config Config
	err := json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func main() {
	config := getConfig()
	fmt.Println(config)
}
