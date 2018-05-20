package main

import "fmt"

func main() {
	buckets := make([]int, 1)
	fmt.Println(buckets[0])
	buckets[0] = 42
	fmt.Println(buckets[0])
	buckets[0]++ // int slice plus plus
	fmt.Println(buckets[0])
}
