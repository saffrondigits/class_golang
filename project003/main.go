package main

import "fmt"

// Swapped numbers
func main() {
	a := 5

	b := 9

	var temp int

	temp = a
	a = b
	b = temp

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}
