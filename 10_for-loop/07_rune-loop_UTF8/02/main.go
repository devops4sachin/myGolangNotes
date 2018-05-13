package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i))) // [] : convert string into a list of byte, we are type casting here.

	}
}
