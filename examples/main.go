package main

import "github.com/mushroomsir/mimetypes"
import "fmt"

func main() {

	fmt.Println(mimetypes.Lookup("x.jpg"))
	// output:  image/jpeg
	fmt.Println(mimetypes.Lookup(".jpg"))
	// output:  image/jpeg
	fmt.Println(mimetypes.Extension("image/jpeg"))
	// output:  jpeg
}
