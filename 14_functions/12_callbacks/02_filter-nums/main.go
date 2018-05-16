package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int { // "callback" is a function name return true for every number greater than 1
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n) // list of all numbers greater than 1
		}

		fmt.Printf("%T", callback) // This will return "func(int) bool"
	}
	return xs // returning slice of int
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs) // prints [2 3 4]
}
