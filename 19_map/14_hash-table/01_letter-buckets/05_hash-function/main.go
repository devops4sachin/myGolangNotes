package main

import "fmt"

func main() {
	n := hashBucket("Go", 12)
	fmt.Println(n)
}

func hashBucket(word string, buckets int) int { // Creating a hash function. Takes word and number of buckets.
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
