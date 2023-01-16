package main

import (
	"fmt"
	"strconv"
)


func main() {
	n, err := strconv.Atoi("42")
	
	fmt.Println("Converted number:", n)
	fmt.Println("Returned error value:", err)
	fmt.Println()
	n, err = strconv.Atoi("hahah")
	fmt.Println("Converted number:", n)
	fmt.Println("Returned error value:", err)
	fmt.Println()
	input := "42"
	n, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Printf("Success: value %q converted to %d!\n", input, n)
}