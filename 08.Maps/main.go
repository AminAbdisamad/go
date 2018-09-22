package main

import (
	"fmt"
)

func main() {
	// Creating maps
	Users := make(map[string]string)
	Users["name"] = "Mohamed Amin"
	Users["Email"] = "Amin@gmail.com"
	Users["Phone"] = "990-44-453"
	Users["Address"] = "here"
	fmt.Println(Users)
}
