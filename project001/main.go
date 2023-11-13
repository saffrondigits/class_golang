package main

import "fmt"

func main() {
	// Declare a variable to store user's input
	var input string

	fmt.Println("Welcome to my small project using Golang!")
	fmt.Println("Enter your name")

	fmt.Scanln(&input)

	fmt.Println("Hello", input)

	var age int
	fmt.Println("How old are you?")
	fmt.Scanln(&age)

	if age < 18 {
		fmt.Println("You are not eligible to get a driving license")
	} else {
		fmt.Println("You can get a driving license")
	}

	fmt.Println("Thank you for using my small project using Golang!")
}
