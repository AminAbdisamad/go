package main

import (
	"fmt"
)

func main() {
	// Creating maps
	Users := make(map[string]string)
	Users["name"] = "Mohamed"
	Users["Email"] = "test@gmail.com"
	Users["Phone"] = "990-44-453"
	Users["Address"] = "here"
	// Display name
	fmt.Println(Users["name"])
	// Legth of users
	fmt.Println(len(Users))

	// Delete map value
	delete(Users, "name")
	Users["name"] = "Hassan"
	fmt.Println(Users["name"])
	fmt.Println(Users)

	// Create map and assign key values
	Customers := map[string]string{
		"Name": "Geedi", "Email": "Geedi@gmail.com", "Phone": "5757-484-453",
	}

	// Looping through map
	for  _,value := range Customers {
		fmt.Println(value)

	}
	// fmt.Println(Customers)
}
