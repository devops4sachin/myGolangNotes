package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57} // We are creating a Slice over here explicitly
	n := average(data...)                     // variadic number of arguments, data is one item and we have to open it further
	fmt.Println(n)
}

func average(sf ...float64) float64 { // For parameters the dots are in front, for arguments the dots are the end
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// Alternative would be as below, in the next program.
// func average(sf []float64) float64 {
