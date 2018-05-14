package main

import "fmt"

func main() {
	greet("Jane", "Doe")
}

func greet(fname string, lname string) { // Two parameters
	fmt.Println(fname, lname)
}
