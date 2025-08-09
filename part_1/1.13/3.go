package main

import "fmt"

func main() {
	var val1 int = 5
	var val2 float32 = 5.32
	fmt.Printf("%.1f\n", float32(val1)) // 5.0,
	// Printf рассматривается в следующем разделе

	fmt.Println(int(val2)) // 5
}
