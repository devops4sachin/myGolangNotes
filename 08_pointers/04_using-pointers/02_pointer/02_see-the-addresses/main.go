package main

import "fmt"

func zero(z *int) {
	fmt.Println(z) // It would be same address.
	*z = 0
}

func main() {
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x) // x is 0
}
