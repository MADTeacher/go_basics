package main

import "fmt"

type employee struct {
	name     string
	string   // первое анонимное поле
	uint8    // второе анонимное поле
	position string
}

func main() {
	emp := employee{
		name:     "Tom",
		position: "Intern",
	}
	emp1 := employee{
		name:     "Alex",
		position: "Intern",
		uint8:    17,
		string:   "R&D",
	}
	// присваиваем значение анонимному полю
	emp.uint8 = 22
	emp.string = "R&D"
	fmt.Printf("%+v\n", emp1)      //{name:Alex string:R&D uint8:17 position:Intern}
	fmt.Printf("%+v\n", emp)       // {name:Tom string:R&D uint8:22 position:Intern}
	fmt.Printf("%+v\n", emp.uint8) // 22 - вывод значения анонимного поля
}
