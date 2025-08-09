package main

import (
	"flag"
	"fmt"
)

func main() {
	// конфигурируем флаги:
	// Первый аргумент - название флага
	// Второй - значение по умолчанию
	// Третий - описание того, что делает флаг
	myStr := flag.String("str", "value", "StrFlag description")
	myInt := flag.Int("int", 0, "IntFlag description")
	myBool := flag.Bool("bool", false, "BoolFlag description")

	// Запускаем парсинг флагов, подаваемых в командной строке
	flag.Parse()

	// используем флаги
	fmt.Println("String flag:", *myStr)
	fmt.Println("Int flag:", *myInt)
	fmt.Println("Bool flag:", *myBool)
}
