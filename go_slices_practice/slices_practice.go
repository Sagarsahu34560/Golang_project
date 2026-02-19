package main

import "fmt"

func main() {
	slices()
}

/*
	Slices are similar to arrays, but are more powerful and flexible.

	Like arrays, slices are also used to store multiple values of the same type in a single variable.

	However, unlike arrays, the length of a slice can grow and shrink as you see fit.

	In Go, there are several ways to create a slice:

	1. Using the []datatype{values} format
	   Syntax: slice_name := []datatype{values}

	2. Create a slice from an array

	3. Using the make() function
*/

func slices() {

	mySlice1 := []int{}                          // this code declares empty slice with 0 length and 0 capacity
	mySlice2 := []int{1, 2, 3, 3, 4, 4, 4, 6, 7} // the code declares slice of integers with length (9) and capacity (9)

	fmt.Println(mySlice1)
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))

	fmt.Println(mySlice2)
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))

}

// func sliceArray() {

// 	var myArray = []int{}
// }
