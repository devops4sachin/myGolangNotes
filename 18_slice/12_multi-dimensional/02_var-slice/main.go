package main

import (
	"fmt"
)

func main() {
	var student []string    // Creating a Slice using the var, will get a nil slice
	var students [][]string // Var will always set to whatever you are declaring to nil/0 value, in this case is it nil
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil) // true
}
