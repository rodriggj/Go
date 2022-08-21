package main

import "fmt"

func main() {
	var s string
	s = "how are you?"
	s = `how are you?`

	fmt.Println(s)

	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)

	s = `
<html>
	<body>"Hello"</body>
</html>
	`
	fmt.Println(s)

	// Use case concerning file path
	fmt.Println("c:\\my\\dir\\file") // need to escape backslashes
	fmt.Println(`c:\my\dir\file`)    // don't need to use escape characters
}
