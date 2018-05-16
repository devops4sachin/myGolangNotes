package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1) // calling recursion
}

func main() {
	fmt.Println(factorial(4))
}
