package main

import "fmt"

func main() {

	myGreeting := map[string]string{}   // empty initialization using composite literal
	myGreeting["Tim"] = "Good morning." // work with composite literal
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
