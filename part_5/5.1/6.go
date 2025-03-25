package main

import "fmt"

type employee struct {
	name           string
	departmentName string
	age            uint8
	position       string
}

func main() {
	emp2 := employee{
		name:     "Tom",
		position: "Intern",
	}
	empPoint := &emp2       // первый способ через оператор &
	empPoint1 := &employee{ // второй способ через оператор &
		name:     "Maxim",
		position: "Intern",
		age:      18,
	}
	// изменение значений полей через указатель на структуру
	(*empPoint).departmentName = "R&D"
	// тоже самое, что
	empPoint.age = 23
	empPoint1.departmentName = "R&D"
	fmt.Printf("%+v\n", *empPoint)  // {name:Tom departmentName:R&D age:23 position:Intern}
	fmt.Printf("%+v\n", *empPoint1) // {name:Maxim departmentName:R&D age:18 position:Intern}

	empPoint2 := new(employee)      // использование ключевого слова new
	fmt.Printf("%+v\n", *empPoint2) // {name: departmentName: age:0 position:}
	empPoint2.departmentName = "Oo"
	empPoint2.name = "Alex"
	empPoint2.position = "Engineer"
	empPoint2.age = 40
	fmt.Printf("%+v\n", *empPoint2) // {name:Alex departmentName:Oo age:40 position:Engineer}
}
