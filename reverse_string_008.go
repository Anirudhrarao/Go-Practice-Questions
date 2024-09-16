package main

import (
	"fmt"
)

func reverseString(s string) string {
	// Convert the string to a slice of runes to handle multi-byte characters
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}

	return string(runes)
}

func main() {
	str := "Ganpati"
	reversed := reverseString(str)
	// Print the original and reversed string
	fmt.Println("Original String:", str)
	fmt.Println("Reversed String:", reversed)
}
