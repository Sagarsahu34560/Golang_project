package main

import "fmt"

func main() {
	clicalculator()
}

func clicalculator() {
	var num1, num2 float64
	var operator string

	// Code for num1
	fmt.Print("Enter first number: \n")
	fmt.Scan(&num1)

	// Code for Operators
	fmt.Print("Enter Operator (+, -, *, /, %)\n")
	fmt.Scanln(&operator)

	// Code for num2
	fmt.Print("Enter the second number: \n")
	fmt.Scanln(&num2)

	// Code for calculating
	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		if num1 == 0 {
			fmt.Printf("Result: Infinity\n")
		} else if num2 == 0 {
			fmt.Printf("Result: Cannot divide by zero\n")
		} else {
			fmt.Printf("Result: %.2f\n", num1/num2)
		}
	case "%":
		fmt.Printf("Result: %.0f\n", float64(int(num1)%int(num2)))
	default:
		fmt.Printf("Invalid Input! Please try again.\n")
	}

}
