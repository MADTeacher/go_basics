package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

func (p *Person) incAge() {
	p.age++
	if p.age > 30 {
		defer func() {
			fmt.Println("Calling defer function")
		}()
		panic(fmt.Sprintf("%s too old!!!", p.name))
		defer func() { // нет смысла в вызове отложенной функции после паники
			fmt.Println("Useless function")
		}()
	}
}

func main() {

	alex := &Person{
		name: "Alex",
		age:  27,
	}
	for i := 0; i < 6; i++ {
		alex.incAge()
	}
	fmt.Printf("Person data: %+v\n", alex)
}
