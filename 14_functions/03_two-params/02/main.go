package main

import "fmt"

func main() {
	greet("Jane", "Doe")
}

func greet(fname, lname string) { // writing fname string, lname string we can write it this way
	fmt.Println(fname, lname)
}
