package main

import (
	"fmt"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	x := 10
	y := 20
	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
	// passing addresses of x and y to the swap
	swap(&x, &y)
	fmt.Printf("After swap: x = %d, y = %d\n", x, y)
}
