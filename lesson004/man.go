package main

import "fmt"

func main() {
	var number int
	var price float32
	var text string

	fmt.Println(&number) // Memory location of integer variable number
	fmt.Println(&price)  // Memory location of float variable price
	fmt.Println(&text)   // Memory location of string variable text
}
