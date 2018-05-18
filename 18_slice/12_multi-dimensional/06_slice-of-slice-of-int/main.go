package main

import (
	"fmt"
)

func main() {

	transactions := make([][]int, 0, 3) // Slice of slice of int

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0, 4) // Slice of int
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j) // Populating the slice of int
		}
		transactions = append(transactions, transaction) // Appending to slice of slice of int
	}
	fmt.Println(transactions)
}
