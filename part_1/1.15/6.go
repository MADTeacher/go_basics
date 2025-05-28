package main

import (
	"fmt"
	"strings"
)

func main() {
	// Разбиваем строку на срезы по любым пробельным символам:
	// пробел, табуляция, перенос строки и т.д.
	str1 := "  Йо   is бедствие  "
	fieldsRes := strings.Fields(str1)
	fmt.Print("\nFields(str1): ")
	fmt.Printf("%q\n", fieldsRes)

	// Разбиваем строку на срезы и задаем
	// пользовательскую функцию для определения разделителя
	str2 := "A+B+++C+DE++F"
	fieldsFuncRes := strings.FieldsFunc(str2, func(r rune) bool {
		return r == '+' // '+' - разделитель
	})
	fmt.Print("\nFieldsFunc(str2, '+'): ")
	fmt.Printf("%q\n", fieldsFuncRes)
}
