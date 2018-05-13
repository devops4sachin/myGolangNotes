package main

import "fmt"

func main() {

	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a) // printing the address of variable "a".
	fmt.Printf("%d \n", &a)
}
