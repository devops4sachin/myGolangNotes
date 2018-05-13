package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break // breaking a for loop, with no condition set above.
		}
		i++
	}
}
