package disgard

import (
	"fmt"
	"path"
)

func disgard() {
	var file string
	_, file = path.Split("css/main.css")
	fmt.Println("file:", file)
}
