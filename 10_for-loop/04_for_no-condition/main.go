package main

import "fmt"

func main() {
	i := 0
	for { // for loop with no condition, https://golang.org/doc/effective_go.html#for
		fmt.Println(i)
		i++
	}
}
