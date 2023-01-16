package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]

	ft, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("error: '%s' is not a number.\n", input)
		return
	}
	div := 3.2808
	meter := float64(ft) / div
	fmt.Printf("%g feet is %.2f meters.\n", ft, meter)


}