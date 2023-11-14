package main

import "fmt"

func main() {
	value1 := 3
	value2 := 4

	total := add(value1, value2)

	fmt.Printf("Total: %v\n", total)

	val1, val2 := returnsTwoValues()

	fmt.Printf("Value 1: %v\n", val1)
	fmt.Printf("Value 2: %v\n", val2)
}

// add function accepts two parameters and return one type
func add(a int, b int) int {

	sum := a + b

	return sum
}

func returnsTwoValues() (int, bool) {
	myInt := 11
	myBool := true

	return myInt, myBool
}
