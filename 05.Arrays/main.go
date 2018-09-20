package main

import (
	"fmt"
)

func main() {
	// Arrays - in Go arrays size should be specified before using it
	var arr [4]string
	arr[0] = "Orange"
	arr[1] = "Blue"
	arr[2] = "Green"
	arr[3] = "Apple"
	Users := [4]string{"Hassan", "Mohamed", "Hussien", "Ali"}
	fmt.Println(Users)

	// Slices - Extended arrays
	people := []string{"Ali", "Haji", "Hassan", "Mohamed", "Hussien"}
	//Length
	fmt.Println(len(people))
	//Range
	fmt.Println(people[1:4])

	// userInfo := []string{"Mohamed Amin"}

	// fmt.Println("%T\n", people)

}
