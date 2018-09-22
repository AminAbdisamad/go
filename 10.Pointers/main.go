package main

import "fmt"

func main() {
	a := 5
	// Pointer
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	// Reading the value of b
	fmt.Println(*b)

	// changing the value of pointers
	*b = 20
	println(a)
}
