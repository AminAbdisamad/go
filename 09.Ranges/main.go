package main

import "fmt"

func main() {
	// Create slice of customers
	CustomerNames := []string{"Hassan", "Ali", "Hussein", "Ilyas", "Kali"}
	// Looping through slices
	for _, CustomerName := range CustomerNames {
		fmt.Printf("Customer Name: %s\n", CustomerName)
	}

	// Creating map
	Users := map[string]string{
		"Name": "Ali", "Email": "Ali@gmail.com",
	}
	for k, v := range Users {
		fmt.Printf("%s : %s\n", k, v)
	}
}
