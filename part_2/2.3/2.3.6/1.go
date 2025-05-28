package main

import (
	"fmt"

	"github.com/jaswdr/faker"
)

func main() {
	faker := faker.New()

	fmt.Println("----------Names----------")
	for i := 0; i < 3; i++ {
		fmt.Println(faker.Person().Name())
	}
	fmt.Println("----------Beer----------")
	for i := 0; i < 3; i++ {
		fmt.Println(faker.Beer().Name())
	}
	fmt.Println("----------Address----------")
	for i := 0; i < 3; i++ {
		fmt.Println(faker.Address().Address())
	}
	fmt.Println("----------Country----------")
	for i := 0; i < 3; i++ {
		fmt.Println(faker.Address().Country())
	}
	fmt.Println("----------Email----------")
	for i := 0; i < 3; i++ {
		fmt.Println(faker.Internet().Email())
	}
	fmt.Println("-----Phone Number--------")
	for i := 0; i < 3; i++ {
		fmt.Println(faker.Phone().Number())
	}
}
