package main

import "fmt"

// Factorial of a numbers
func main() {
	var num int

	fmt.Printf("Enter a number: ")
	fmt.Scanln(&num)

	if num < 0 {
		fmt.Println("Invalid input! Please enter a positive integer")
		return
	}

	result := Factorial(num)

	fmt.Printf("Factorial of %v is %v!\n", num, result)

}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	fact := 1

	for i := 1; i <= n; i++ {
		fact = fact * i
	}

	return fact
}
