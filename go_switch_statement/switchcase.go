package main

import "fmt"

func main() {
	switchcase()
}

func switchcase() {

	day := 4

	switch day {
	case 1: // Single-Case Switch Case
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3, 4, 5: // Multi-case Switch Case
		fmt.Println("Wednesday")
	default:
		fmt.Println("Holiday")
	}

}
