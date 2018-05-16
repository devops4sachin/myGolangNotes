package main

import "fmt"

func visit(numbers []int, callback func(int)) { // name of function is "callback" and the type is "func(int)", numbers is slice of int
	for _, n := range numbers { // We are passing a function as an argument
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) { // we are passing an anonymous function here
		fmt.Println(n)
	})
}

// callback: passing a func as an argument
