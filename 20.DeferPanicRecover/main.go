package main

/****
Summary

Defer
- used to delay the execution of statement until function exits
- useful to group "open and close" function together
- run LIFO - last-in first-out
- Arguments evaluated at the time defer is executed, not time of called function executed

Panic
- Ocur when application can't continue at all

-Recover
- used to recover from panic
- only useful in defered functions







****/

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func getRobotsFromGoogle(){
	// Defer := Executes any function thats passed into it after function finishes its final statement but before it returns
	// Defer statement doesn't take function itself it takes function call

	res, err := http.Get("https://www.google.com/robots.txt")
	fmt.Println(res.Body)
	if err !=nil{
		log.Fatal(err)
	}
	defer res.Body.Close() // close after main function is executed
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(robots)
}

func serveInternet(){
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello world"))
		log.Println("Server started on port 8080 ...")
	})
	err:= http.ListenAndServe(":8080",nil)
	if err != nil{
		panic(err.Error())
	}
}

func panicker(){
	fmt.Println("Its about to panic")
	defer func(){
		if err:= recover(); err != nil{
			log.Println("Error: ",err)
		}
	}()
	panic("Something went wrong")
	fmt.Println("Stop panicking")
}
func main(){

	// fmt.Println("start")
	// panicker()
	// fmt.Println("end")

	// Panics happen after defer statement is executed

	serveInternet()
}