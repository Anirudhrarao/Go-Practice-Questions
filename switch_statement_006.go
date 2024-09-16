package main

import (
	"fmt"
)

func main() {
	day := 3
	// switch
	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}
}
