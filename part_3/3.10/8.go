package main

import "fmt"

func GetMapValues[K comparable, V any](m map[K]V) []V {
	// вернет срез из значений словаря
	values := make([]V, len(m))
	idx := 0
	for _, v := range m {
		values[idx] = v
		idx++
	}
	return values
}

func GetMapKeys[K comparable, V any](m map[K]V) []K {
	// вернет срез из ключей словаря
	keys := make([]K, len(m))
	idx := 0
	for k, _ := range m {
		keys[idx] = k
		idx++
	}
	return keys
}

func CreateNewMap[K comparable, V any](k []K, v []V) map[K]V {
	// создание словаря из двух срезов
	newMap := make(map[K]V)
	if len(k) < len(v) {
		for idx, elem := range k {
			newMap[elem] = v[idx]
		}
	} else {
		for idx, elem := range v {
			newMap[k[idx]] = elem
		}
	}
	return newMap
}

func main() {
	intSlice := []int{1, 3, 4, 6, 7}
	stringSlice := []string{"Oo", "^_^", "-_-"}
	myMap := map[int]string{
		1:   "Alex",
		2:   "Maxim",
		200: "Jon",
	}
	fmt.Println(GetMapKeys(myMap))                   // [1 2 200]
	fmt.Println(GetMapValues(myMap))                 // [Alex Maxim Jon]
	fmt.Println(CreateNewMap(intSlice, stringSlice)) // map[1:Oo 3:^_^ 4:-_-]
}
