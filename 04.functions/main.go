package main

import (
	"fmt"
	"strings"
)

func greeting(name string)  {
	fmt.Println("Hello ", name)
}

func addTwoN(a, b int) int {
	return a + b
}
// Taking function as a param
func greatNames(names []string, f func(name string)){
	for _,value := range names {
		f(value) // Invoking the function here
	}

}

// Return Multiple values
func getInitials(n string)(string,string){
	// get initials, capitalize and return
	s := strings.ToUpper(n)
	namesList := strings.Split(s, " ")
	var initials = []string{}
	for _, value:=range namesList {
		initials = append(initials,value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0],"_"


}

func main() {
	// fmt.Println(greeting("Aminux"))
	fmt.Println(addTwoN(4, 8))

	var names = []string{"Adi","Ali","Has","Kaz"}
	greatNames(names, greeting);

	//
	fN, lN := getInitials("Hassan Ahmed")
	fmt.Println(fN, lN)
}


