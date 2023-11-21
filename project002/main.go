package main

import "fmt"

// TODO: Reverse a number
func main() {
	number := -123

	reverse := 0
	for number != 0 {
		rem := number % 10
		reverse = reverse*10 + rem
		number = number / 10
	}

	fmt.Println(reverse)
}
