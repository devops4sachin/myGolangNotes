package main

import "fmt"

func main() {

	if true { // true is a reserved word here, in Python its True
		fmt.Println("This ran")
	}

	if false {
		fmt.Println("This did not run")
	}
}
