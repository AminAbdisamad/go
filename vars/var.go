package main

import "fmt"

// GO - is case-sensative
func main() {
	// Using var - using var declaration outside the function is possible
	// var firstName string = "Hassan"
	var firstName, lastName = "Hassan", "Amin"

	// short way
	company := "Loopysec"
	age := 29

	//Constant variable
	const nickname string = "Aminux"
	var isMarried = true
	var GPA = 3.75

	//Shorter way to assign variables
	address, email, phone := "Serdivan", "aminux@gmail.com", "555-666-777"
	fmt.Println(firstName, lastName, nickname, company, age, address, email, phone, isMarried, GPA)

	// Find variable type
	fmt.Printf("%T\n", GPA)
}
