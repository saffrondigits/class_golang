package main

import "fmt"

func main() {
	admin := "dave"

	var inputUser string
	fmt.Scanln(&inputUser)

	if inputUser == admin {
		fmt.Println("Admin user is trying to login")
	} else {
		fmt.Println("General user is trying to login")
	}
}
