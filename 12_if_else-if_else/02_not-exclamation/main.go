package main

import "fmt"

func main() {

	if !true { // Negation of True, is False
		fmt.Println("This did not run")
	}

	if !false {
		fmt.Println("This ran")
	}

}
