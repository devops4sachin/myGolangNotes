package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus":
		fmt.Println("Wassup Marcus") // This would be executed.
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi") // This would be also executed.
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian") // This would be also executed.
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}
}
