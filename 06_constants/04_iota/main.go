package main

import "fmt"

const (
	a = iota // 0, second way to declare iota
	b        // 1
	c        // 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
