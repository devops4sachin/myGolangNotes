package main

import "fmt"

func wrapper() func() int { // The return here is "func() int" and not only int. Thus wrapper returns the whole function.
	x := 0
	return func() int { // This function body is returned actually.
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
