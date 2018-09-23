package main

import (
	"fmt"
	"strconv"
)

func main() {
	customer1 := Customers{firstName: "hassan", lastName: "Amin", country: "Somalia", gender: "M", address: "Godey", phone: "555-667-433", age: 20}
	fmt.Println(customer1)
	fmt.Println(customer1.firstName)
	customer1.firstName = "Ali"
	fmt.Println(customer1.firstName)
	fmt.Println(customer1.greeting())

}

// strucks are more like classes in other PL.
type Customers struct {
	firstName, lastName, gender, address, phone, country string
	age                                                  int
}

// Methods - we 2 kinds 1- value receivers and pointer receiver

// Method - value reciever

func (c Customers) greeting() string {
	return "Hi, My full Name is " + c.firstName + c.lastName + "I'm " + strconv.Itoa(c.age) + " Years old from " + c.country
}
