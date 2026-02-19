package main

import "fmt"

func main() {
	challenge1()
	challenge2()
	challenge3()
}

// Print numbers 1 to 100 divisible by 3
func challenge1() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

// Take user marks and print grade using switch
func challenge2() {
	var marks int
	fmt.Println("Enter your marks (0-100):")
	fmt.Scanln(&marks)

	switch {
	case marks >= 90 && marks <= 100:
		fmt.Println("A")
	case marks >= 75:
		fmt.Println("B")
	case marks >= 65:
		fmt.Println("C")
	case marks >= 50:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
}

// Print multiplication table for a number
func challenge3() {
	var num int
	fmt.Println("Enter the number to get the multiplication table: ")
	fmt.Scan(&num)

	fmt.Printf("The multiplication table is of %d: \n", num)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", num, i, num*i)
	}
}
