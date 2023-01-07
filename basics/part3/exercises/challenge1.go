package main
import (
	"fmt"
	"os"
	"strings"
)
// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}
	username := "jack"
	password := "4444"
	if strings.Compare(os.Args[1], username) == 0 {
		if strings.Compare(os.Args[2], password) == 0 {
			fmt.Printf("Access granted to %q\n", os.Args[1])
		} else {
			fmt.Printf("Invalid password for %q\n", os.Args[1])
		}
	} else {
		fmt.Printf("Access denied for %q\n", os.Args[1])
	}
}