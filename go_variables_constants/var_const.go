package main

import "fmt"

// ------------VARIALES---------------

func main() {
	showVariable()
	varDecWithoutInitialValue()
	decMultipleVariable()
	varDecInBlock()
	constDeclaration()

}

func showVariable() {

	// Long form [Syntax: var variableName = value]
	var name string = "Sagar" // type is string

	// Type inerence (Go figures out the type)
	var surname = "Sahu" // type is inferred

	// Short declaration (only used inside function) [Syntax: var := value]
	age := 27 // type is inferred

	// Constant declaration [Syntax: const CONSTNAME type = value]
	/* const keyword declare the variable as constant,
	because it means unchangable and read-only*/
	const country = "India" // type is inferred

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Surname", surname)
	fmt.Println("Country:", country)
}

// Declaring a variable without initial values
func varDecWithoutInitialValue() {
	var a string
	var b int
	var c bool

	fmt.Println(a) // " "
	fmt.Println(b) // 0
	fmt.Println(c) // false
}

// Decalring multiple variable
func decMultipleVariable() {

	// Using the Type keyword in the variable declaration
	var a, b, c, d int = 1, 2, 3, 4

	// not using the Type keyword in the variable declaration
	var q, w = 1, "Hello"
	r, t := 7, "world"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(q)
	fmt.Println(w)
	fmt.Println(r)
	fmt.Println(t)
}

func varDecInBlock() {

	// Variable is being declared inside a block
	var (
		a int
		b int    = 1
		c string = "Halloween"
	)

	d := "I am here" // Short declaration is not allowed inside the var(...) block

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

// Note: for PUBLIC variable we use capital letter at first while defining the variable

var PublicVariableName string = "this is the public variable"

//------------CONASTANT-----------------

const PI = 3.14

// Two types of constant

// Typed Constant
const A int = 1

// Untyped Constant
const B = 2

// Multiple contant declaration
const D, E = 5, "Hello"

// Constant declaration inside a block
const (
	T int    = 20 // In constant, it is must to declare a value in this variable
	P int    = 10
	L string = "Dare You"
)

func constDeclaration() {
	fmt.Println(PI)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(D)
	fmt.Println(E)

	fmt.Println(T)
	fmt.Println(P)
	fmt.Println(L)
}
