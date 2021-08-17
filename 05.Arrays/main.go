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
	arr[3] = "Red"
	fmt.Println(arr)

	Users := [4]string{"Hassan", "Mohamed", "Hussien", "Ali"}
	fmt.Println(Users)

	// Slices - Extended arrays
	people := []string{"Ali", "Haji", "Hassan", "Mohamed", "Hussien"}

	//Length
	fmt.Println(len(people))
	//Slice Range
	fmt.Println(people[1:4])

	// userInfo := []string{"Mohamed Amin"}

	fmt.Printf("%T\n", people)

	// Appending elements to slices
	people = append(people, "Gorad")
	for i:=0; i<len(people); i++{
		fmt.Println("Your name is ", people[i])
	}

	// Create empty slice
	var cityNames []string;
	// Add a single city
	cityNames = append(cityNames, "Sakarya")
	fmt.Println(cityNames)
	cityNames = append(cityNames, "Istanbul")
	fmt.Println(cityNames)


}
