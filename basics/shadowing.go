package main

import "fmt"

var i = 10

// variable shadowing occurs when a variable declared within 
// a certain scope such as decision block, method, or inner class 
// has the same name as a variable declared in an outer scope

func main() {
	fmt.Println("package's i:", i)

	// package's i is being shadowed:
	i := 5
	// i above is a new variable.
	
	fmt.Println("main block:", i)

	// a new scope begins
	{
		// i is a new variable
		i := 6
		{
			i := 2
			fmt.Println("double inner block:", i)
		}
		fmt.Println("inner block:", i)
	}
	// the scope ends

	// main's i
	fmt.Println("main's i:", i)
}