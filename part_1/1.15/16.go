package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder

	// Write - записывает байты (строку) в Builder
	n, err := b.Write([]byte("Hello, "))
	if err != nil { // проверка на успешное завершение операции
		panic(err)
	}
	fmt.Printf("After Write: %q (bytes written: %d)\n", b.String(), n)

	// WriteString - записывает строку
	n, _ = b.WriteString("world")
	fmt.Printf("After WriteString: %q (bytes written: %d)\n", b.String(), n)

	// WriteByte - записывает один байт (символ)
	b.WriteByte('!')
	fmt.Printf("After WriteByte: %q\n", b.String())

	// WriteRune - записывает один Unicode-символ
	n, _ = b.WriteRune('🌍')
	fmt.Printf("After WriteRune: %q (bytes written: %d)\n", b.String(), n)

	// Len - возвращает текущую длину строки
	fmt.Printf("Current length: %d\n", b.Len())

	// Cap - возвращает текущую емкость внутреннего буфера
	fmt.Printf("Current capacity: %d\n", b.Cap())

	// Grow - увеличивает емкость буфера структуры Builder
	b.Grow(100)
	fmt.Printf("After Grow(100), new capacity: %d\n", b.Cap())

	// Reset - очищает содержимое буфера Builder-а
	b.Reset()
	fmt.Printf("After Reset: %q, length: %d\n", b.String(), b.Len())
}
