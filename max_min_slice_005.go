package main

import (
	"fmt"
	"math"
)

func findMaxMin(nums []int) (int, int) {
	if len(nums) == 0 {
		fmt.Println("The slice is empty.")
		return 0, 0
	}
	// Initialize min and max with extreme values
	min := math.MaxInt64
	max := math.MinInt64

	// Iterate through the slice to find max and min
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}

func main() {
	nums := []int{23, 42, 17, 68, 5, 99, -2, 37}
	max, min := findMaxMin(nums)
	fmt.Printf("Maximum: %d, Minimum: %d\n", max, min)
}
