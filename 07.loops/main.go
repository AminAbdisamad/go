package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Number %d\n", i)
	// }
	fmt.Println(fizzBuzz(100))
	
       // The init and post statements are optional.	
	sum := 1
	for ; sum < 1000; { sum += sum }
	fmt.Println(sum)
	
	// For is Go's "while"
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	// Forever If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
	for { fmt.Println("Invinite")
}
}

func fizzBuzz(n int) int {
	// if n / 3 = 0 if n/5 = 0
	for i := 1; i < n; i++ {

		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}

	}
	return 0
}
