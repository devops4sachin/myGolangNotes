package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!", // remember we have trailing comma here
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
