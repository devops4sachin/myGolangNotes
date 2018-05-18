package main

import (
	"fmt"
)

func main() {
	student := []string{}
	students := [][]string{}
	student[0] = "Todd" // Will get error here, as we didn't set the lenght here, the underlying array is of lenght 0
	// student = append(student, "Todd") // This should work.
	fmt.Println(student)
	fmt.Println(students)
}
