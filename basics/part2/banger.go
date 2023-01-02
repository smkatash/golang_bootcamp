package main

import (
	"fmt"
	"os"
	"strings"
)

// Usage go run banger.go Hello
func main() {
	msg := os.Args[1]
	l := len(msg)
	r := strings.Repeat("!", l)
	s := r + msg + r
	s = strings.ToUpper(s)
	fmt.Println(s)
}