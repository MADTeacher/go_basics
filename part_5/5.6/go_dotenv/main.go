package main

import (
	"fmt"
	"go_dotenv/dotenv"
	"log"
)

func main() {
	env := dotenv.NewDotEnv(false)

	err := env.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	adminID, err := env.GetInt("ADMIN_ID")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(adminID)

	adminPassword, err := env.GetString("ADMIN_PASSWORD")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(adminPassword)

	adminMode, err := env.GetBool("ADMIN_MODE")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(adminMode)

	adminEmail, err := env.GetString("ADMIN_EMAIL")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(adminEmail)
}
