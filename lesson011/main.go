package main

import "fmt"

func main() {
	printComputerConf()
	greeting()
	printComputerConf()
}

func greeting() {
	fmt.Println("Hello World!")
}

func printComputerConf() {
	fmt.Println("Brand Name: MSI")
	fmt.Println("Product Name: GF99")
	fmt.Println("Processor: i9")

	keyBoard()
}

func keyBoard() {
	fmt.Println("Brand: HP")
	fmt.Println("Type: QWERTY")
}
