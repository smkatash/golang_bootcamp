package main
import "fmt"

func main() {
	var (
		speed int
		heat float64
		off bool
		brand string
	)
	// T -type
	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)
	fmt.Println(heat, off, brand)

	var (
		planet = "venus"
		distance = 261
		orbital = 224.7701
		hasLife = false
	)
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v mln kms\n", distance)
	fmt.Printf("Orbital Area: %v days\n", orbital)
	fmt.Printf("Planet has life: %v\n", hasLife)
	fmt.Println()
	// *** Type safety ***
	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d mln kms\n", distance)
	fmt.Printf("Orbital Area: %f days\n", orbital)
	fmt.Printf("Planet has life: %t\n", hasLife)
	
	fmt.Printf("%v is %v away. Think! %[2]v kms! %[1]v OMG!\n", planet, distance)
	
	fmt.Println()
	fmt.Printf("Orbital Area: %f days\n", orbital)
	fmt.Printf("Orbital Area: %.0f days\n", orbital)
	fmt.Printf("Orbital Area: %.1f days\n", orbital)
	fmt.Printf("Orbital Area: %.4f days\n", orbital)
}