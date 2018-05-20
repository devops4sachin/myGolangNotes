package main

import "fmt"

func main() {
	word := "Hello"
	letter := rune(word[0]) // Convert H to rune, means to int32 alias
	fmt.Println(letter)
}
