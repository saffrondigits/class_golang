package main

import "fmt"

func main() {
	fmt.Printf("Hello %v, how are you doing?\n", "Ravi")
	Sum(4, 0)
	fmt.Printf("%v, %v, %v are the students who passed the exam\n", "Ravi", "KD", "Bob")
}

func Sum(a int, b int) {
	if a == 0 {
		fmt.Printf("total sum is: %v\n", b)
		return
	}

	if b == 0 {
		fmt.Printf("total sum is: %v\n", a)
		return
	}

	fmt.Printf("total sum is: %v\n", a+b)
}
