package main

import (
	"fmt"
	"time"
)

func main() {
	practiceChallenge1()
	practiceChallenge11()
	practiceChallenge2()
	practiceChallenge3()
}

// 1. Create a program that calculates your birth year from your age.
func practiceChallenge1() {
	var age int
	const currentYear = 2025

	fmt.Println("Please enter your age")
	fmt.Scanln(&age)

	birthYear := currentYear - age

	fmt.Printf("Your birth year is: %d.\n", birthYear)
}

// using time package
func practiceChallenge11() {
	var age int
	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)

	now := time.Now()
	currentYear := now.Year() // gets the current date and time from the system clock
	birthYear := currentYear - age

	fmt.Printf("Your birth year is : %d.\n", birthYear)
	fmt.Printf("Your age is %d, and born in %d, and today date is %d, %s %d.\n", age, birthYear, now.Year(), now.Month(), now.Day())

}

// 2. Create 3 constants: company, role, and experience — print them.
func practiceChallenge2() {
	const (
		company    = "TCS"
		role       = "Software Developer"
		experience = 3.5
	)

	fmt.Printf("I work in %s as %s with an experience of %v years.\n", company, role, experience)
}

// 3. Create variables a and b, and print sum, diff, product, and quotient.
func practiceChallenge3() {
	var a, b int
	fmt.Println("Give a number a: ")
	fmt.Scan(&a)
	fmt.Println("Give a number b: ")
	fmt.Scan(&b)

	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b

	fmt.Println(sum, diff, product, quotient)

	// 4. Use fmt.Printf to print values with proper format specifiers.
	fmt.Printf(" sum = %d,\n diff = %d,\n product = %d,\n quotient = %d\n", sum, diff, product, quotient)
}
