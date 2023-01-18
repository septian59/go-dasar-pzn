package main

import "fmt"

func main() {
	var name1 = "ian"
	var name2 = "budi"

	var result bool = name1 == name2

	fmt.Println(result) // hasil false

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)  // false
	fmt.Println(value1 < value2)  // true
	fmt.Println(value1 == value2) // false
	fmt.Println(value1 != value2) // true

}
