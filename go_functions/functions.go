package main

import "fmt"

//################################################################
// // 1. Basic function
// /*
// Syntax:
// func FunctionName() {
// 	code to be executed
// }
// */
func hello() {
	fmt.Println("Hey there!!")
}

func main() { // A function is called inside 'main()'
	hello()
	hello() // A function can be called multiple times
}

//################################################################

// // 2. Parameters and Arguements
// /*
// Syntax:
// func FunctionName(param1 type, param2 type, param3 type) {
// 	// code to be executed
// }
// */
// func classmates(name string, group int) { // Here 'name' and 'group' are the parameters
// 	fmt.Println(name, "is my classmate and is from group", group)
// }

// func main() {
// 	classmates("chikki", 1) // Here values are the arguements eg: (chikki and 1)
// 	classmates("sagar", 2)  //  Likewise we can have multiple parameters inside the function
// 	classmates("dj", 3)
// 	classmates("mona", 4)
// }

//################################################################

// // 3. Function with return values
// /*
// Syntax:
// func FunctionName(param1 type, param2 type) type {
//   // code to be executed
//   return output
// }
// */

// func returnFunction(x int, y int) int {
// 	return x + y
// }
// /* Here, myFunction() receives two integers (x and y)
//    and returns their addition (x + y) as integer (int) */
// func main() {
// 	fmt.Println(returnFunction(1, 2))
// }

// // 4. Named return values
// func namedReturnValues(a int, b int) (result int) {
// 	result = a + b
// 	return
// }

// func main() {
// 	fmt.Println(namedReturnValues(10, 2))
// }

// // 5. Store the return value in a variable
// func storeReturnValue(a int, b int) (result int) {
// 	result = a + b
// 	return
// }

// func main() {
// 	total := storeReturnValue(1, 5)
// 	fmt.Println(total)
// }

// // 6. Multiple return functions
// func multipleFunctions(a string, b int) (resultNum int, resultText string) {
// 	resultNum = b + b
// 	resultText = a + " You!"
// 	return
// }

// func main() {
// 	fmt.Println(multipleFunctions("I love", 1500))
// }

// // 7. Here we can store the two return values in two different variables
// func returnInTwoVariable(c string, a int, b string) (resultInt int, resultString string, resultString1 string) {
// 	resultInt = 1000 * a
// 	resultString = b + " Ironman!"
// 	resultString1 = "I love " + c
// 	return
// }

// func main() {
// 	x, y, z := returnInTwoVariable("you!", 3, "I am")
// 	fmt.Println(z, x, y)
// }

// // 8. Omitting a return value
// func addMultiply(a int, b int) (int, int) {
// 	sum := a + b
// 	product := a * b

// 	return sum, product
// }

// func main() {
// 	s, _ := addMultiply(4, 5) // Here we have omitted the p value with '_'
// 	fmt.Println("sum:", s)

// 	_, p := addMultiply(4, 5) // Here we have omitted the p value with '_'
// 	fmt.Println("product:", p)
// }


// ################################################################


// // Example 1
// func addMultiply(a int, b int) (int, int) {
// 	sum := a + b
// 	product := a * b

// 	return sum, product
// }

// func main() {
// 	s, p := addMultiply(4, 5)
// 	fmt.Println("sum:", s, "product:", p)
// }


