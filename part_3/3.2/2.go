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

func (r RUB) setNewValue(rub RUB) {
	r = rub
	fmt.Printf("В кошельке теперь %d рублей\n", r)
}

func main() {
	var rub RUB = 3475
	rub.setNewValue(9000)
	fmt.Printf("В кошельке на самом деле %d рублей\n", rub)
}
