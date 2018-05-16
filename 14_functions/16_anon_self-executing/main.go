package main

import "fmt"

func main() {
	func() { // Anonymous self executing function
		fmt.Println("I'm driving!")
	}() // This is for self execution over here itself.
}
