package main
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// if n, err := strconv.Atoi("42") ; err == nil {
	// 	fmt.Println("Successful, n is", n)
	// }

	if a := os.Args; len(a) != 2 {
		fmt.Println("Provide a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Cannot convert %q.\n", a[1])
	} else {
		fmt.Printf("%s * 2 = %d.\n", a[1], n * 2)
	}
}