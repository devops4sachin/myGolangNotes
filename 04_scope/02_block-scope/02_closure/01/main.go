package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{ // you can have a closure with in a closure.
		fmt.Println(x)
		y := "The credit belongs with the one who is in the ring."
		fmt.Println(y)
	}
	// fmt.Println(y) // outside scope of y, will not print it.
}
