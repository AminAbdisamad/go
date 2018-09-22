package main

import (
	"fmt"
)

func main() {
	sum := add()
	for i := 0; i < 50; i++ {
		fmt.Println(sum(i))
	}
}

// Addition function

func add() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
