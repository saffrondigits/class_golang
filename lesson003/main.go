package main

import "fmt"

func main() {
	var balance float32
	var balanceAcc2 float32
	var total float32

	balance = 100.00

	fmt.Println(balance)

	balance = balance + 50

	fmt.Println(balance)

	balanceAcc2 = 350.55

	total = balance + balanceAcc2

	fmt.Println(total)
}

// Default value of int is 0
// Default value of float is 0
// Default value of string is ""
// Default value of bool is false
// Default value of byte is 0
