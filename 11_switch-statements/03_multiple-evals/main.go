package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny": // Multiple conditions, Tim or Jenny.
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")
	}
}
