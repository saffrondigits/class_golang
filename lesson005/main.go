package main

import "fmt"

func main() {
	var noOfUnit int

	fmt.Println("Enter number of units:")
	fmt.Scanln(&noOfUnit)

	fmt.Println("Total units", noOfUnit)
	fmt.Println("A String", 5, true, 99, 33.5)
}
