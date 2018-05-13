package main

import "fmt"

func main() {
	for i := 250; i <= 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := 'a'
	bar := "a"
	fmt.Println(foo)               // prints 97, numeric value for rune.
	fmt.Printf("%T \n", foo)       // prints int32
	fmt.Printf("%T \n", rune(foo)) // prints int32
	fmt.Printf("%T \n", rune(bar)) // this will error as, cannot convert bar(type string) to type rune.
}

/*
NOTE:
Some operating systems (Windows) might not print characters where i < 256

If you have this issue, you can use this code:

fmt.Println(i, " - ", string(i), " - ", []int32(string(i)))

UTF-8 is the text coding scheme used by Go.

UTF-8 works with 1 - 4 bytes.

A byte is 8 bits.

[]byte deals with bytes, that is, only 1 byte (8 bits) at a time.

[]int32 allows us to store the value of 4 bytes, that is, 4 bytes * 8 bits per byte = 32 bits.
*/
