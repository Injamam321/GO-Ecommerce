package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "i am imjamam hossain. i am junior software engineer.")

}

func main() {
	mux := http.NewServeMux() // router

	mux.HandleFunc("/hello", helloHandler) //route

	mux.HandleFunc("/about", aboutHandler) //route

	fmt.Println("Server running on: 3000")

	err := http.ListenAndServe(":3000", mux) //"Failed to start the server"

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
