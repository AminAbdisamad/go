package main

import (
	"fmt"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func confirm_age(age int) string {
if age < 20 {return ("You can't be here")}
return ("Old enough to do whatever you want :) ")
}

func main() {
	
	
 fmt.Println(sqrt(2), sqrt(-4))
 age := confirm_age(27)
 fmt.Println(age)

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
