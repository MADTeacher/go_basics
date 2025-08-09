package main

import (
	"fmt"
	"strings"
)

func main() {
	str := " !my string! "

	// удяляем пробелы с обеих сторон
	fmt.Println(strings.TrimSpace(str)) // !my string!

	// удяляем заданные символы в начале и в конце строки
	fmt.Println(strings.Trim(str, " "))  // !my string!
	fmt.Println(strings.Trim(str, " !")) // my string

	// удяляем заданные символы в начале строки
	fmt.Println(strings.TrimLeft(str, " "))   // !my string!
	fmt.Println(strings.TrimLeft(str, " m!")) // y string!

	// удяляем заданные символы в конце строки
	fmt.Println(strings.TrimRight(str, " "))    // !my string!
	fmt.Println(strings.TrimRight(str, "n!g ")) //  !my stri

	// удяляем заданный символ или подстроку в начале строки
	fmt.Println(strings.TrimPrefix(str, " !my")) // string!

	// удяляем заданный символ или подстроку в конце строки
	fmt.Println(strings.TrimSuffix(str, "g! ")) // !my strin
}
