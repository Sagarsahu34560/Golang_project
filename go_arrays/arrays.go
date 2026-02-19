package main

import "fmt"

func main() {
	arrays()
	arrayElementAccess()
	arrayElementChange()
	arrayInitialization()
	arrayLength()
}

func arrays() {
	fmt.Println()

	// Here length is defined

	/* Using "var" keyword
	Syntax: var array_name = [length] datatype {values}

	Using ":=" symbol
	Syntax: array_name := [length] datatype {values} */

	var arr1 = [4]int{1, 2, 3, 4}
	arr2 := [5]int{5, 6, 7, 8}
	var arrString1 = [3]string{"bus", "car", "auto"}

	fmt.Println(arr1, arr2, arrString1)

	// Here length is inferred

	/* Using "var" keyword
	Syntax: var array_name = [...] datatype {values}

	Using ":=" symbol
	Syntax: array_name := [...] datatype {values} */

	var arr3 = [...]int{1, 2, 3, 4}
	arr4 := [...]int{5, 6, 7, 8, 9, 10}
	var arrString2 = [...]string{"bus", "car", "auto"}

	fmt.Println(arr3, arr4, arrString2)

}

// Access elements of an array
func arrayElementAccess() {

	prices1 := [4]int{10, 20, 30, 40}
	fmt.Println(prices1[2])
}

// Change elements of an array
func arrayElementChange() {

	prices2 := [5]string{"ten", "twenty", "thirty", "forty"}
	fmt.Println(prices2)

	prices2[4] = "fifty"
	fmt.Println(prices2)
}

// Array Initialization
func arrayInitialization() {

	arr5 := [5]int{}              // not initialized
	arr6 := [5]int{1, 2, 3}       // partially initialized
	arr7 := [5]int{1, 2, 3, 4, 5} // fully initialized

	fmt.Println(arr5, arr6, arr7)

	// Initializing Specific Elements
	arr8 := [5]int{0: 20, 3: 46}

	fmt.Println(arr8)
}

// Array Length
func arrayLength() {

	arr9 := [5]string{"me", "my", "mine", "more", "memory"}
	arr10 := [...]int{1, 2, 3, 4, 5, 5, 6, 666}

	fmt.Println(arr9, len(arr9))
	fmt.Println(arr10, len(arr10))
}
