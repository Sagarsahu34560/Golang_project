package main

import "fmt"

// ------------OUTPUT---------------

func main() {
	output()
}

func output() {

	var a = 15.5
	var text = "Hello world"

	// Output Functions

	fmt.Print(a)
	fmt.Println(text) // Println is used to add the next line at the end of the arguements
	fmt.Print(a, text)
	fmt.Print(a, "/n", text) // "/n" is used to shift the next arguement to the next line

	// Printf() function and General Formatting Verbs
	fmt.Printf("in this line the value: %v and the type: %T", a, a)
	fmt.Printf("in this line the text: %v and the type: %T", text, text)

	// Some more types of general formatting verb
	fmt.Printf(" this prints the value in go-syntax format %#v", text)
	fmt.Printf("this print the modulo sign, eg: %v%%", a)
}

/* There are more types of Formatting verb

----------------------------------------------------------------------------

## For simple understnding

Verb |  Meaning
----------------
%d	 |  Integer
%f	 |  Float
%s	 |  String
%t	 |  Boolean
%v	 |  Any value (auto-format)

-----------------------------------------------------------------------------
## For detailed explanation
-----------------------------------------------------------------------------

1. Integer Formating verbs

Verb	Description
----    -----------
%b	    Base 2
%d	    Base 10
%+d	    Base 10 and always show sign
%o	    Base 8
%O	    Base 8, with leading 0o
%x	    Base 16, lowercase
%X	    Base 16, uppercase
%#x	    Base 16, with leading 0x
%4d	    Pad with spaces (width 4, right justified)
%-4d	Pad with spaces (width 4, left justified)
%04d	Pad with zeroes (width 4)

-----------------------------------------------------------------------------

2. String formatting verbs

Verb	Description
----    -----------
%s	    Prints the value as plain string
%q	    Prints the value as a double-quoted string
%8s	    Prints the value as plain string (width 8, right justified)
%-8s	Prints the value as plain string (width 8, left justified)
%x	    Prints the value as hex dump of byte values
% x	    Prints the value as hex dump with spaces

-----------------------------------------------------------------------------

3. Boolean Formatting verbs

Verb	Description
----    -----------
%t	    Value of the boolean operator in true or false format (same as using %v)

-----------------------------------------------------------------------------

4. Float Formatting verbs

Verb  | Description
----  | -----------
%e	  | Scientific notation with 'e' as exponent
%f	  | Decimal point, no exponent
%.2f  | Default width, precision 2
%6.2f | Width 6, precision 2
%g	  | Exponent as needed, only necessary digits

-----------------------------------------------------------------------------

*/
