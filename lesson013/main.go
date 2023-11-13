package main

import "fmt"

func main() {
	add(4, 3)

	add(5, 5)

	name := "Ravi"
	Greet(name)
}

func add(a int, b int) {
	var sum int
	sum = a + b

	fmt.Println("Total is:", sum)
}

func Greet(name string) {
	fmt.Println("Hello", name)
}
