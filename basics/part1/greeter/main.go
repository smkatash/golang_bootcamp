package main
import ( "fmt"; "os" )


func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st:", os.Args[1])
	fmt.Println("2nd:", os.Args[2])
	fmt.Println("3rd:", os.Args[3])
	fmt.Println("Length:", len(os.Args))
}