package main

import "fmt"

func main() {
	number := 2

	if number == 5 {
		fmt.Println("The number is 5")
	} else if number > 5 {
		fmt.Println("The number greater then 5")
	} else {
		fmt.Println("The number lesser then 5")
	}
}
