package main

import (
    "fmt"
    "reflect"
    "strconv"
)

func main() {
    val1 := 33.321
    str1 := strconv.FormatFloat(val1, 'f', –1, 64)
    // равносильно
    str2 := fmt.Sprintf("%.3f", val1)
    fmt.Println(reflect.TypeOf(str2) , str2) // string 33.321
    fmt.Println(reflect.TypeOf(str1), str1) // string 33.321 
    val2, _ := strconv.ParseFloat(str2, 64)
    fmt.Println(reflect.TypeOf(val2), val2) // float64 33.321
}
