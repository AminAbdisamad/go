package main

import (
	"fmt"
	"strings"
	"sort"
)

func main (){
	//! 1. strings Package
	var message string = "Hello there, I love go programming"
	// fmt.Println(strings.c(message, "Hello"))
	if strings.Contains(message,"there"){
		fmt.Printf("Your string %v contains %q \n", message,"there" )
	}

	a,b := "message 1", "message 2";
	if strings.Compare(a,b) == 0 {
		fmt.Println("a = ", a, "b = ", b)
	}else{
		fmt.Println("Value is different")
	}
	// fmt.Println(strings.Compare(a,b))

	// ! 2. Sorts package

	var ages []int = []int{12,45,11,56,67,2}
	// var ages = []int{12,45,11,56,67,2}
	// var ages = []int{12,45,11,56,67,2}
	// Sorts all numbers
	sort.Ints(ages)

	fmt.Println(ages)

	index:= sort.SearchInts(ages,450)

	if (index != len(ages)) {
		fmt.Println("Found", index)
	}

	names := []string{"hassan", "abdi","mustaf","faysal"}
	sort.Strings(names) // alters the orignal value
	fmt.Println(names)

	value:= sort.SearchStrings(names, "mustafaa")
	fmt.Println(value)

}