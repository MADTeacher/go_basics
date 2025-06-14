package main

import (
	"fmt"

	"go_json_store/jsonstore"
)

func main() {
	store := jsonstore.NewJSONStore("store.json")
	store.SetValue("map", map[string]int{"a": 1, "b": 2})
	store.SetValue("strList", []string{"a", "b", "c"})
	store.SetValue("intList", []int{1, 2, 3})

	// Получение и приведение типов
	rawStrList := store.GetList("strList")
	var strList []string
	for _, v := range rawStrList {
		if s, ok := v.(string); ok {
			strList = append(strList, s)
		}
	}

	rawIntList := store.GetList("intList")
	var intList []int
	for _, v := range rawIntList {
		if n, ok := v.(int); ok {
			intList = append(intList, n)
		}
	}

	rawMap := store.GetMap("map")
	myMap := make(map[string]int)
	for k, v := range rawMap {
		switch n := v.(type) {
		case int:
			myMap[k] = n
		case float64:
			myMap[k] = int(n)
		}
	}

	fmt.Printf("%T\n", strList) // []string
	fmt.Printf("%T\n", intList) // []int
	fmt.Printf("%T\n", myMap)   // map[string]int

	fmt.Println(strList) // [a b c]
	fmt.Println(intList) // [1 2 3]
	fmt.Println(myMap)   // map[a:1 b:2]

	boolVal := store.GetBool("bool")
	fmt.Println(*boolVal) // true
	intVal := store.GetInt("int")
	fmt.Println(*intVal) // 55
	floatVal := store.GetFloat("double")
	fmt.Println(*floatVal) // 99.4
}

//***************  3  ***************************
/*
func main() {
	store := jsonstore.NewJSONStore("store.json")

	store.ResetValue("map")
	store.ResetValue("str")

	fmt.Println(store.GetValue("map")) // <nil>
	fmt.Println(store.GetValue("str")) // <nil>
}
*/

//***************  2  ***************************
/*
func main() {
	store := jsonstore.NewJSONStore("store.json")

	store.SetValue("strList", "-_-")
	store.SetValue("double", 99)

	fmt.Println(store.GetValue("strList")) // -_-
	fmt.Println(store.GetValue("double"))  // 99
}
*/

//***************  1  ***************************
/*
func main() {
	store := jsonstore.NewJSONStore("store.json")
	store.SetValue("strList", []string{"a", "b", "c"})
	store.SetValue("int", 55)
	store.SetValue("bool", true)
	store.SetValue("double", 3.14)
	store.SetValue("map", map[string]int{"a": 1, "b": 2})
	store.SetValue("str", "(づ˶•༝•˶)づ♡")

	fmt.Println(store.Values())
	// [[a b c] 55 true 3.14 map[a:1 b:2] (づ˶•༝•˶)づ♡]

	fmt.Println(store.Keys())
	// [bool double map str strList int]

	fmt.Println(store.Contains("strList")) // true
	fmt.Println(store.GetValue("strList")) // [a b c]
	fmt.Println(store.GetValue("int"))     // 55
	fmt.Println(store.GetValue("bool"))    // true
	fmt.Println(store.GetValue("double"))  // 3.14
	fmt.Println(store.GetValue("map"))     // map[a:1 b:2]
	fmt.Println(store.GetValue("str"))     // (づ˶•༝•˶)づ♡
} */

/*func main() {
	fmt.Println("Starting JSON Store Demo")
	store := jsonstore.NewJSONStore("store.json")

	// Set some values
	mapValue := map[string]any{"a": 1, "b": 2}
	strList := []any{"a", "b", "c"}
	intList := []any{1, 2, 3}

	store.SetValue("map", mapValue)
	store.SetValue("strList", strList)
	store.SetValue("intList", intList)

	// Get values
	retrievedStrList := store.GetList("strList")
	retrievedIntList := store.GetList("intList")
	retrievedMap := store.GetMap("map")

	// Print results
	fmt.Printf("strList type: %T\n", retrievedStrList)
	fmt.Printf("intList type: %T\n", retrievedIntList)
	fmt.Printf("map type: %T\n", retrievedMap)

	fmt.Printf("strList: %v\n", retrievedStrList)
	fmt.Printf("intList: %v\n", retrievedIntList)
	fmt.Printf("map: %v\n", retrievedMap)

	// Let's demonstrate more features similar to the commented section in the Dart code
	store.SetValue("int", 55)
	store.SetValue("bool", true)
	store.SetValue("double", 3.14)
	store.SetValue("str", "(づ˶•༝•˶)づ♡")

	fmt.Println("\nAll values:", store.Values())
	fmt.Println("All keys:", store.Keys())

	fmt.Printf("Contains 'strList': %v\n", store.Contains("strList"))
	fmt.Printf("Value of 'strList': %v\n", store.GetValue("strList"))
	fmt.Printf("Value of 'int': %v\n", store.GetValue("int"))
	fmt.Printf("Value of 'bool': %v\n", store.GetValue("bool"))
	fmt.Printf("Value of 'double': %v\n", store.GetValue("double"))
	fmt.Printf("Value of 'map': %v\n", store.GetValue("map"))
	fmt.Printf("Value of 'str': %v\n", store.GetValue("str"))

	// Update values
	store.SetValue("strList", "-_-")
	store.SetValue("double", 99)
	fmt.Printf("\nUpdated 'strList': %v\n", store.GetValue("strList"))
	fmt.Printf("Updated 'double': %v\n", store.GetValue("double"))

	// Reset values
	store.ResetValue("map")
	store.ResetValue("str")
	fmt.Printf("After reset, 'map': %v\n", store.GetValue("map"))
	fmt.Printf("After reset, 'str': %v\n", store.GetValue("str"))
}
*/
