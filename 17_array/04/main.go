package main

import "fmt"

func main() {
	var x [256]byte

	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = byte(i) // 8 bit, but it will show as "uint8"
	}
	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v) // it will show the type as "uint8"
		if i > 50 {
			break
		}
	}
}
