package main

import (
	"fmt"
)

func main() {
	fmt.Println(greeting("Aminux"))
}
func greeting(name string) string {
	return "Hello" + name
}
