package main

import "fmt"

func main() {
	mySlice := make([]int, 1) // setting lenght and capacity to 1
	fmt.Println(mySlice[0])   // prints 0
	mySlice[0] = 7
	fmt.Println(mySlice[0]) // prints 7
	mySlice[0]++            // increment 7 by 1, to 8
	fmt.Println(mySlice[0]) // prints 8
}
