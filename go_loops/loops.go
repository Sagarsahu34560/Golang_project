package main

import "fmt"

func main() {
	forloop()
	whilelikeloop()
	infiniteloop()
	continuestatement()
	breakstatement()
	rangekeyword()
	nestedloop()
}

/*
1. for loop
2. while like loop
3. infinite loop
4. continue statement
5. break statement
6. range keyword
7. idx(index) and val(value)
8. nested loop
*/

func forloop() {
	for i := 1; i <= 8; i++ {
		fmt.Println("count:", i)
	}

	for j := 0; j <= 100; j += 10 {
		fmt.Println(j)
	}
}

func whilelikeloop() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func infiniteloop() {
	for {
		fmt.Println("print...")
		break
	}
}

func continuestatement() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		if i == 2 {
			continue
		}
	}
}

func breakstatement() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		if i == 2 {
			break
		}
	}
}

// Syntax: for index, value := range array/slice/map
func rangekeyword() {
	fruits := [4]string{"apple", "orange", "banana", "jackfruit"}
	for index, value := range fruits {
		fmt.Printf("index: %v\tvalue: %v\n", index, value) // here '\t' - space, '\n' - next line
	}

	// Here we can omit index and only show the value
	for _, value := range fruits {
		fmt.Printf("%v\n", value)
	}

	// Here we can omit the value and only show the index
	for index, _ := range fruits {
		fmt.Printf("%v\n", index)
	}
}

func nestedloop() {
	adj := [2]string{"good", "bad"}
	fruits := [4]string{"boy", "girl", "man", "lady"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
}
