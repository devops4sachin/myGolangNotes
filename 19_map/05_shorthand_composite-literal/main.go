package main

import "fmt"

func main() {

	myGreeting := map[string]string{ // Non empty composite literal initialization, composed of literal values
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	fmt.Println(myGreeting["Jenny"])
}
