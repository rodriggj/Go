package main

import (
	"fmt"
	"path"
)

func main() {
	// *******Example 1: Both Return values Defined******
	// var dir string
	// var file string

	// dir, file = path.Split("css/main.css")

	// fmt.Println("dir: ", dir)
	// fmt.Println("file:", file)

	// *******Example 2: Disgarding a Return Value******
	// var file string
	// _, file = path.Split("css/main.css")
	// fmt.Println("file:", file)

	// *******Example 3: Short Declaration******
	_, file := path.Split("css/main.css")
	fmt.Println("file: ", file)
}
