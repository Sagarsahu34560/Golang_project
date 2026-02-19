package main

import "fmt"

func main() {
	personalInfo()
}

func personalInfo() {

	var name string
	var age int
	const country = "India"

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Printf("\nHello %s!\n", name)
	fmt.Printf("Your age is %d and you belong to %s and we call you %s", age, country, name)
}
