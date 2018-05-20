package main

import "fmt"

func main() {
	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12) // number, letter, hash value
	}
}
