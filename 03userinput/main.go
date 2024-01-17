package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	// comma ok

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter anything: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Result: ", input)
	fmt.Printf("Type: %T", input)
}
