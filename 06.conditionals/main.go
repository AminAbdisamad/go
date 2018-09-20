package main

import (
	"fmt"
)

func main() {

	//IF ELSE
	color := "who"
	if color == "red" {
		fmt.Println("The color is red")

	} else if color == "blue" {
		fmt.Println("The color is blue")

	} else {
		fmt.Println("The color is NOT blue Nor red")

	}

	// Switch
	switch color {
	case "red":
		println("Color is red")
	case "blue":
		println("Color is blue")
	case "green":
		println("Color is green")
	case "white":
		println("Color is white")
	case "black":
		println("Color is black")
	case "orange":
		println("Color is orange")
	default:
		println("I don't know this color")

	}

}
