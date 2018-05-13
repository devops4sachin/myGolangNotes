package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	max := max(7)    // This is a variable shadowing, declaring the variable with same name "max", not recommended.
	fmt.Println(max) // max is now the result, not the function
}

// don't do this; bad coding practice to shadow variables
