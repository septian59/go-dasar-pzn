package main

import "fmt"

func main() {
	const firstName string = "septian"
	const lastName = "primansyah"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	//multiple variable
	const (
		age     int = 0
		address     = "ketapang"
	)

	fmt.Println(age)
	fmt.Println(address)
}
