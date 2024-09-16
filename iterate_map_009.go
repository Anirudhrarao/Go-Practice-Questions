package main

import (
	"fmt"
)

func main() {
	fruits := map[string]int{
		"Apple":  5,
		"Banana": 12,
		"Orange": 9,
	}

	// Iterate through the map
	for fruit, qty := range fruits {
		fmt.Printf("Fruits: %s, Quantity: %d\n", fruit, qty)
	}
}
