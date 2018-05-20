package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	// var myGreeting map[string]string  ## This will create a nil reference, in this case you cannot use the "append" function
	myGreeting["Tim"] = "Good morning." // For nil map this would give an error.
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
