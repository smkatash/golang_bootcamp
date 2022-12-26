package main
import ( "fmt"; "path" )

func main() {
	var dir, file string
	var photo string

	dir, file = path.Split("css/main.css")
	_, photo = path.Split("folder/photo.jpg")
	_, document := path.Split(("fd/doc.pdf"))

	fmt.Println(dir, " ", file)
	fmt.Println(photo)
	fmt.Println(document)
}