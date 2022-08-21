package main

import (
	"fmt"
)

func main() {
	speed := 5   //speed is currently typed as an int
	force := 2.6 //force is currenlty typeda as float64

	speed = int(float64(speed) * force)

	fmt.Println("speed:", speed)
}
