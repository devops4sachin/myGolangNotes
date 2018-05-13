package main

import "fmt"

var x = 42 // This is a package level scope. "x" is accessbile in the whole package.
// Its accessible to all the files within the same package.

func main() {
	fmt.Println(x) // 'x' is accessible here as well.
	foo()
}

func foo() {
	fmt.Println(x)
}
