package main

import "fmt"

func main() {
	var i int

	for i = 0; i > 10; i++ {
		fmt.Println("Value of i", i)
	}
}

// for initialization; condition check; increment/decrement {

// }

// 1. Initialization happens only once
// 2. Condition check happens multiple times before for loop block gets executed
// 3. Increment/Decrement happens at the end of the loop block
