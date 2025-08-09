package main

import "fmt"

type MyStruct[K string | uint, T string | []string] struct {
	id   K
	data T
}

func (m *MyStruct[K, T]) getID() K {
	return m.id
}

func (m *MyStruct[K, T]) getData() T {
	return m.data
}

func (m *MyStruct[K, T]) setData(data T) {
	m.data = data
}

func main() {
	myStruct := MyStruct[uint, string]{
		id:   1,
		data: "Oo",
	}
	fmt.Printf("K = int, T = string:  %+v\n", myStruct)
	myStruct.setData("^_^")
	fmt.Printf("K = int, T = string:  %+v\n", myStruct)

	newStruct := MyStruct[string, []string]{
		id:   "0xA321",
		data: []string{"10", "20", "30"},
	}
	fmt.Printf("K = string, T = []string:  %+v\n", newStruct)
	newStruct.setData([]string{})
	fmt.Printf("K = string, T = []string:  %+v\n", newStruct)
}
