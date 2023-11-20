package main

import "fmt"

func main() {
	var countOfApple int

	for i := 0; i < 5; i++ {
		countOfApple = countOfApple + 2
	}

	fmt.Printf("Total apples are: %v\n", countOfApple)
}

// for initialization; condition check; increment/decrement {

// }

// 1. Initialization happens only once
// 2. Condition check happens multiple times before for loop block gets executed
// 3. Increment/Decrement happens at the end of the loop block
