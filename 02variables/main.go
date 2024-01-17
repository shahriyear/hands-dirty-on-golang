package main

import "fmt"

var LoginToken string = "abcd"

func main() {
	var name string = "Fahim"
	fmt.Println(name)
	fmt.Printf("Variable type is %T \n", name)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type is %T \n", isLoggedIn)

	var number int = 123
	fmt.Println(number)
	fmt.Printf("Variable type is %T \n", number)

	var number2 int32 = 123
	fmt.Println(number2)
	fmt.Printf("Variable type is %T \n", number2)

	var number3 float32 = 123.123
	fmt.Println(number3)
	fmt.Printf("Variable type is %T \n", number3)

	var number4 uint16 = 123
	fmt.Println(number4)
	fmt.Printf("Variable type is %T \n", number4)

	number5 := 123.99
	fmt.Println(number5)
	fmt.Printf("Variable type is %T \n", number5)

	var test = 123.99
	fmt.Println(test)
	fmt.Printf("Variable type is %T \n", test)

	fmt.Println(LoginToken)
	fmt.Printf("Variable type is %T \n", LoginToken)
}
