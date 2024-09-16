package main

import (
	"fmt"
)

func sumMapValues(data map[string]int) int {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum
}

func main() {
	data := map[string]int{
		"apple":  10,
		"banana": 20,
		"orange": 15,
	}

	total := sumMapValues(data)
	fmt.Printf("The total sum of the map values is: %d\n", total)
}
