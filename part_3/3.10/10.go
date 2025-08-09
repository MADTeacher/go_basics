package main

import "fmt"

func CreateNewMap[K comparable, V any](k []K, v *[]V) *map[K]V {
	// создание словаря из среза и указателя на срез
	newMap := make(map[K]V)
	if len(k) < len(*v) {
		for idx, elem := range k {
			newMap[elem] = (*v)[idx]
		}
	} else {
		for idx, elem := range *v {
			newMap[k[idx]] = elem
		}
	}
	return &newMap
}

func main() {
	intSlice := []int{1, 3, 4, 6, 7}
	stringSlice := []string{"Oo", "^_^", "-_-"}
	fmt.Println(CreateNewMap(intSlice, &stringSlice)) // &map[1:Oo 3:^_^ 4:-_-]
}
