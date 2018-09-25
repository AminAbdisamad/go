package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Printf("What's your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Welcome to Golang %s!\n", name)

}
