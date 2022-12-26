package main
import "fmt"

func main() {
	// Long declaration
	var score int

	var (
		// related:
		video string

		// closely related:
		duration int
		current int
	)
	score = 100 ; video = "Tutorial" ; duration = 15 ; current = 2
	fmt.Println(score, video, duration, current)

	// Short declaration
	width, height := 100, 100
	width, colour := 50, "red"

	fmt.Println(width, height, colour)
}