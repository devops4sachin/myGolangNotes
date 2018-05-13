package main

import "fmt"

func main() {
	i := 0
	for i < 10 { // loop with only condition, https://golang.org/doc/effective_go.html#for
		fmt.Println(i)
		i++
	}
}
