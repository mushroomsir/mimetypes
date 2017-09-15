package main

import "github.com/mushroomsir/mimetypes"
import "fmt"

func main() {

	fmt.Println(mimetypes.Lookup(".jpg"))
	// output:  image/jpeg
}
