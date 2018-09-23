package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Starting Server ...")
	http.ListenAndServe(":8000", nil)
}

// Index page
func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Home Page")
}

// About page
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}
