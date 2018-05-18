package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 3) // Type of array, initial length of Slice, Underlying capacity of Slice

	fmt.Println("-----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("-----------------")

	for i := 0; i < 80; i++ { // Capacity is increased as an when required
		mySlice = append(mySlice, i) // append to slice
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice[i])
	}
}
