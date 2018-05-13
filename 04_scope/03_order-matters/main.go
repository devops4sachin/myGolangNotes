package main

import "fmt"

func main() {
	fmt.Println(x) // "x" is not declared. Order matters.
	fmt.Println(y) // but "y" is accessible here, "y" is a package level variable and is in the outer scope.
	x := 42
}

var y = 42
