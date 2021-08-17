package main

import "fmt"


func updateName (n string){
	 n = "Hussein"
}

func updateGeoL(l map[string]float64){
	l["long"] = 40.0
}
func main (){
	// Pass By Value

	// Group A: Non-Pointer Values  : Strings,Ints,Floats,Arrays,Bools,Structs
	// In Group A :  Every time we pass a variable to a function go creates copy of the variable
	name := "Hassan"
	fmt.Println(name)
	updateName(name)
	fmt.Println(name)

	// Group B:  Pointer Wrapper Values: Slices, Maps , Functions

	geoLocation := map[string]float64{
		"lat":745.45,
		"long":0.0,
	}

	fmt.Println(geoLocation)
	updateGeoL(geoLocation)
	fmt.Println(geoLocation)




}