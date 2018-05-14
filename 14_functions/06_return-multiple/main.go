package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe "))
}

func greet(fname, lname string) (string, string) { // Returing multiple values, seperated by comma's
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
