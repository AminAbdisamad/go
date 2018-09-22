package main

import "fmt"

func main() {
	// Create slice of customers
	CustomerNames := []string{"Hassan", "Ali", "Hussein", "Ilyas", "Kali"}
	for i, CustomerName := range CustomerNames {
		fmt.Printf("%d Customer Name: %s", i, CustomerName)
	}

}
