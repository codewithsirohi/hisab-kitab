package controllers

import (
	"fmt"
	"net/http"
)

// Define your route handlers
func HomeController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Custom-Header", "Hello, World, how are you!")
	fmt.Fprintln(w, "Welcome to the home page!")
}

func AboutController(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("NAME"))
	fmt.Fprintln(w, "This is the about page.")
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Up and Running")
}
