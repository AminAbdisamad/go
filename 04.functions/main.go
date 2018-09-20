package main

import (
	"fmt"
)

func main() {
	fmt.Println(greeting("Aminux"))
	fmt.Println(addTwoN(4, 8))
}

func greeting(name string) string {
	return "Hello " + name
}

func addTwoN(a, b int) int {
	return a + b
}
