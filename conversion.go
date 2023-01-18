package main

import "fmt"

func main() {

	// konversi interger
	var nilai32 int32 = 1000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//konversi string
	var name = "septian primansyah"
	var s byte = name[0]
	var sString = string(s)

	fmt.Println(name)
	fmt.Println(sString)
}
