package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world() // This would be differed till the function "main" exits, before the exit of "main" this function is called
	hello()       // prints "hello world"
}

/* A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement,
reached the end of its function body, or because the corresponding goroutine is panickin*/
