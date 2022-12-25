package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello, ") ; fmt.Println("world!")
	fmt.Println(runtime.NumCPU())
}