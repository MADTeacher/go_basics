package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

func (p *Person) getName() string {
	return p.name
}

func (p *Person) getAge() uint8 {
	return p.age
}

func (p *Person) incAge() {
	p.age++
}

func main() {
	alex := &Person{
		name: "Alex",
		age:  26,
	}
	fmt.Printf("Person data: %+v\n", alex)
	alex.incAge()
	fmt.Printf("Person data: %+v\n", alex)
	alex = nil
	alex.incAge()
}
