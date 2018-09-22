package main

import (
	"fmt"
)

func main() {
	customer1 := Customers{firstName: "hassan", lastName: "Amin", gender: "M", address: "Godey", phone: "555-667-433", age: 20}
	fmt.Println(customer1)
	fmt.Println(customer1.firstName)
	customer1.firstName = "Ali"
	fmt.Println(customer1.firstName)
	fmt.Println(customer1.greeting())
	a := A{i: 2, j: 3}

	fmt.Println(a.sumFields(10))
}

// strucks are more like classes in other PL.
type Customers struct {
	firstName, lastName, gender, address, phone string
	age                                         int
}

// Methods - we 2 kinds 1- value receivers and pointer receiver

// Method - value reciever

func (c Customers) greeting() string {
	return "Hi, " + c.firstName + " Welcome to Golang"
}

type A struct {
	i int
	j int
}

func (a A) sumFields(k int) int {
	return a.i + a.j + k
}
