package main

import "fmt"

func find1(slice *[]int, value int, check *bool) {
	*check = false
	for i := range *slice {
		if (*slice)[i] == value {
			*check = true
			break
		}
	}
}

func find2(slice *[]int, value int) bool {
	for i := range *slice {
		if (*slice)[i] == value {
			return true
		}
	}
	return false
}

func main() {
	slice := []int{2, 4, 5, 7, 103, 55}
	var check bool
	value := 2
	find1(&slice, value, &check)
	fmt.Printf("%d contains in slice? %t\n", value, check)
	value = 22
	fmt.Printf("%d contains in slice? %t\n", value, find2(&slice, value))
}

// 2 contains in slice? true
// 22 contains in slice? false
