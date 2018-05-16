package main

import "fmt"

func main() {

	if true && false { // Evaluates to false
		fmt.Println("This did not run")
	}

}
