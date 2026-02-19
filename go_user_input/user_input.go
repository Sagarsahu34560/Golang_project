package main

import "fmt"

func main() {
	userInput()
}

func userInput() {

	var i, j int

	// Scan()
	fmt.Print("Type a number: ")
	fmt.Scan(&i, &j)
	fmt.Println("There is a number: ", i, "and", j)

	// Note: &i, &j are memory locations for i and j

	// Scanln()
	fmt.Print("Type a number: ")
	fmt.Scanln(&i, &j)
	fmt.Println("There is a number: ", i, "and", j)

	// Scanf()
	fmt.Print("Type a number: ")
	fmt.Scanf("%v %v", &i, &j)
	fmt.Println("There is a number: ", i, "and", j)

	// Note: In this example, the inputs are received in exactly
	//       the same way defined in Scanf() formatting.

	// Note: %v tells Scanf() to store the value of the inputs in the variables.
}
