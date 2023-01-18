package main
import (
	"fmt"
	"strconv"
	"os"
	"math"
)

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {
	var (
		year int
		err error
	)
	
	arg := os.Args
	if len(arg) != 2 {
		fmt.Println("Give me a year number")
		return
	} else if year, err = strconv.Atoi(arg[1]); err != nil {
		fmt.Printf("%q is not a valid year\n", arg[1])
		return
	}

	if math.Remainder(float64(year), 4) == 0 {
		if math.Remainder(float64(year), 100) == 0 {
			if math.Remainder(float64(year), 400) == 0 {
				fmt.Printf("%d is a leap year\n", year)
			} else {
					fmt.Printf("%d is not a leap year\n", year)
			}
		} else {
			fmt.Printf("%d is a leap year\n", year)
		}
	} else {
		fmt.Printf("%d is not a leap year\n", year)
	}
}