package main

import (
	"fmt"
)

func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "Todd" // will not get any error here
	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
