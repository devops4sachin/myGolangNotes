package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{ // 1. Block level scope
		fmt.Println(x) // x is accessible here
		y := "The credit belongs with the one who is in the ring."
		fmt.Println(y)
	}
	fmt.Println(y) // y is outside of the scope and is not accessible here
}
