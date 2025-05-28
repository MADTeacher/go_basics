package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hi, my dear friend!!!"
	newStr := strings.Replace(str, "d", "&", 1)
	fmt.Println(newStr) // Hi, my &ear friend!!!

	// замена всех вхождений символа
	fmt.Println(strings.ReplaceAll(str, " ", "_"))
	// Hi,_my_dear_friend!!!

	// замена подстроки
	fmt.Println(strings.ReplaceAll(str, "dear", "good"))
	// Hi, my good friend!!!

	// удаление подстроки или символа
	fmt.Println(strings.ReplaceAll(str, "dear", ""))
	// Hi, my  friend!!!
}
