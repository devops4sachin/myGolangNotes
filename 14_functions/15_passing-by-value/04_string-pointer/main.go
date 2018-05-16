package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println(&name) // 0x82023c080

	changeMe(&name)

	fmt.Println(&name) //0x82023c080
	fmt.Println(name)  //Now prints Rocky
}

func changeMe(z *string) {
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Todd
	*z = "Rocky"    // We are assigning a new string here, "Strings" are immutable in go.
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Rocky
}
