package main

import "fmt"

func main (){

	// 1. fmt.Print -> Prints without newline
	// 2. fmt.Prinln -> Prints with a new line
	// 3. fmt.Printf -> is used for formatted strings
	// see examples
	firstName, lastName := "Hassan", "Ahmed";
	age := 30;
	var price float64 = 45.25;
	// ? Print Function is similar to Println except it doesn't add a new line
	// 1.Printing without newline
	fmt.Print("Hello, ")
	fmt.Print("World\n")

	//2.  Printing with newline
	fmt.Println("My First Name is ", firstName, " And My Last Name is ", lastName)

	//3. FORMATED STRINGS
	// %v - default format for variables
	// %q - quoto inside string variables "quit"
	// %T - returns type of the variable
	fmt.Printf("My name is %v and I'm %v years old %v \n", firstName, age,price)
	fmt.Printf("My name is %q and I'm %q years old \n", firstName, age)
	fmt.Printf("My name is %T and I'm %T years old %T \n", firstName, age,price)

	fmt.Printf("The price is %f USD per unit\n", price)
	// Limit number of zeros for when printing float variable
	fmt.Printf("The price is %0.2f USD per unit\n", price)
}
