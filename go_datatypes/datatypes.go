package main

import "fmt"

func main() {
	datatypes()
	boolean()
	integerSigned()
	integerUnsigned()
	floatDatatype()
	datatypeString()
}

func datatypes() {

	var a bool = true       // Boolean
	var b int = 10          // Integer type
	var c float32 = 25.637  // Floating point number
	var d string = "Rodger" // String

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

/* List of all the datatypes present in Golang

Datatype	      |   Example	   |   Description
------------------|----------------|------------------------
int, int8, int32, |   10	       |   Whole numbers
int64             |                |
float32, float64  |   10.25	       |   Decimal numbers
string	          |   "Hello"	   |   Text
bool	          |   true/false   |   Boolean logic
rune	          |   'A'	       |   Single Unicode character


*/

//##################################################################

// Boolean
func boolean() {
	var b1 bool = true // true
	var b2 = true      // true
	var b3 bool        // false by default
	b4 := true         // true

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)

}

//##################################################################

// Integer
// Two types of int -
func integerSigned() {

	// 1. Signed Integer - it is declared as "int" keyword
	var a int = 500
	var b int = -4500

	fmt.Printf("Type: %T, value: %v\n", a, a)
	fmt.Printf("Type: %T, value: %v\n", b, b)
}

/* More types of Signed Integer

Type	Size	                              Range
--------------------------------------------------------------------------------------------
int	    Depends on platform:                  -2147483648 to 2147483647 in 32 bit systems
        32 bits in 32 bit systems and         and -9223372036854775808 to 9223372036854775807
        64 bit in 64 bit systems	          in 64 bit systems
int8	8 bits/1 byte	                      -128 to 127
int16	16 bits/2 byte	                      -32768 to 32767
int32	32 bits/4 byte	                      -2147483648 to 2147483647
int64	64 bits/8 byte	                      -9223372036854775808 to 9223372036854775807

*/

// 2. Unsigned Integer - It is declared as "uint" keyword
func integerUnsigned() {
	var x uint = 80
	var y uint = 400

	fmt.Printf("type: %T, value: %v\n", x, x)
	fmt.Printf("type: %T, value: %v\n", y, y)
}

/* More types of unsigned Integer

Type	Size	                          Range
-------------------------------------------------------------------------------------
uint	Depends on platform:              0 to 4294967295 in 32 bit systems and
        32 bits in 32 bit systems and     0 to 18446744073709551615 in 64 bit systems
        64 bit in 64 bit systems
uint8	8 bits/1 byte	                  0 to 255
uint16	16 bits/2 byte	                  0 to 65535
uint32	32 bits/4 byte	                  0 to 4294967295
uint64	64 bits/8 byte	                  0 to 18446744073709551615

*/

// Note: If we do not specify the integer type, then the default integer is "int"

//##################################################################

//Float
// The float datatype has two keyword - float32, float64

func floatDatatype() {

	//float32
	var x float32 = 123.78
	var y float32 = 3.4e+38

	//float64
	var r float64 = 1.7e+308

	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)
	fmt.Printf("Type: %T, value: %v\n", r, r)
}

/*

Type	   Size	      Range
--------------------------------------------
float32	   32 bits	  -3.4e+38 to 3.4e+38.
float64	   64 bits	  -1.7e+308 to +1.7e+308.

*/

//##################################################################

// String
// String datatype is used to store the sequence of characters(text)

func datatypeString() {
	var text1 string = "Hi"
	var text2 = "How are you?"
	text3 := "I am fine thank you"
	var text4 string // default value is empty string ""

	fmt.Printf("type: %T, value: %v\n", text1, text1)
	fmt.Printf("type: %T, value: %v\n", text2, text2)
	fmt.Printf("type: %T, value: %v\n", text3, text3)
	fmt.Printf("type: %T, value: %v\n", text4, text4)
}
