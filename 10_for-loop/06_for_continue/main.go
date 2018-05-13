package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue // skip the loop by one iteration.
		}
		fmt.Println(i) // will print only all the odd numbers.
		if i >= 50 {
			break
		}
	}
}
