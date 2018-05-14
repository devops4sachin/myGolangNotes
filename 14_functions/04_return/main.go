package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname) // String print, we are not using Println
}
