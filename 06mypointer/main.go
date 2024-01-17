package main

import "fmt"

func main() {

	// var ptr *int
	// fmt.Println("Value of ptr", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Reference of pointer is", ptr)
	fmt.Println("Value of pointer is", *ptr)

	*ptr = *ptr + 2
	fmt.Println("Sum of pointer is", myNumber)

}
