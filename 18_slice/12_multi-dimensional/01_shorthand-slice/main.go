package main

import (
	"fmt"
)

func main() {
	student := []string{} // Creating a slice using, Short hand way to create a Slice. You will have to use append to add to the slice
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil) // false, this is not nil. The underlying array is of lenght 0
}
