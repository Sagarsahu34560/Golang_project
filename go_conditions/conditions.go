package main

import "fmt"

func main() {
	ifAndElseCondition()
	elseifCondition()
	nestedif()
	operatorsAndConditions()
}

// Control Flow: it means deciding which code runs depending on the conditions
/*
	if age >= 18 -> print 'Adult'
	else -> print 'minor'
*/

func ifAndElseCondition() {
	age := 17

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}
}

func elseifCondition() {
	age := 30

	if age <= 14 {
		fmt.Println("child")
	} else if age <= 20 {
		fmt.Println("teenager")
	} else {
		fmt.Println("adult")
	}
}

func nestedif() {
	num := 30

	if num >= 10 {
		fmt.Println("num is greater than 10")
		if num >= 20 {
			fmt.Println("num is even greater than 20")
		}
	}
}

// Logical Operators -> &&(AND), !(NOT), ||(OR)
func operatorsAndConditions() {
	mark := 480
	if mark >= 400 && mark <= 500 {
		fmt.Println("A")
	} else if mark >= 250 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
