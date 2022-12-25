package main
import "fmt"

func main() {
	speed := 100
	force := 2.5

	d_speed := float64(speed) * force
	speed = int(float64(speed) * force)
	fmt.Println(d_speed, speed)

	var hello float32
	hello = 2.3344444444444
	fmt.Println(int(hello), hello)
}