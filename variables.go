package main

import "fmt"

func main() {

	// strings
	var nameOne string = "Name One"
	var nameTwo = "Name Two"
	var nameThree string
	fmt.Println(nameOne, ",", nameTwo, ",", nameThree)

	nameOne = "mario"
	nameThree = "luigi"
	fmt.Println(nameOne,",", nameTwo,",", nameThree)
	
	nameFour := "Name Four"
	fmt.Println(nameOne,",", nameTwo,",", nameThree, ",", nameFour)

	// ints
	var one int = 1
	var two = 2
	three := 3
	fmt.Println(one, two, three)
	
	// bits and memory
	var tmp1 int8 = 25
	var tmp2 int8 = -128
	var tmp3 uint8 = 255
	fmt.Println(tmp1, tmp2, tmp3)
	
	// float
	var scoreOne float32 = 25.93
	var scoreTwo float64 = 25933747448488484848.01
	scoreThree := 3445.01
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}