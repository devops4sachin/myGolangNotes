package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35) // Creating a slice using make, preferred way to make the slices
	students := make([][]string, 35)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
