package main

import "fmt"

type RUB uint // пользовательский именованный тип RUB

const RUB2USD uint = 61
const RUB2EUR uint = 65

//связываем тип RUB с методом конвертации рублей в доллары
func (r RUB) convertRUB2USD() uint {
	return uint(r) / RUB2USD
}

//связываем тип RUB с методом конвертации рублей в евро
func (r RUB) convertRUB2EUR() uint {
	return uint(r) / RUB2EUR
}

func main() {
	var rub RUB = 3475
	fmt.Printf("%d рублей примерно = %d долларов\n", rub, rub.convertRUB2USD())
	fmt.Printf("%d рублей примерно = %d евро\n", rub, rub.convertRUB2EUR())
}
