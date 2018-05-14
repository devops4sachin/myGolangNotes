package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b { // if 'b' is true then assign food = Chocolate, at the same time we are initialising.
		fmt.Println(food)
	}

}
