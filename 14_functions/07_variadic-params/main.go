package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 { // This is variadic function, which accepts Zero or more arguments
	fmt.Println(sf)         // printing sf, it prints a Slice
	fmt.Printf("%T \n", sf) // asking for type, which is []float64
	var total float64
	for _, v := range sf { // ignoring the indexes of Slice
		total += v
	}
	return total / float64(len(sf)) // len of Slice is a function
}
