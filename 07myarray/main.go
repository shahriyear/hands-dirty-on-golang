package main

import "fmt"

func main() {

	var colors [4]string

	colors[0] = "red"
	colors[1] = "green"
	colors[3] = "blue"
	fmt.Println("Colors: ", colors)
	fmt.Println("Colors: ", len(colors))

	var fruits = [4]string{"apple", "orange", "banana"}
	fmt.Println("Fruits", fruits)
	fmt.Println("Fruits", len(fruits))
}
