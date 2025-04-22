package main

import (
	"fmt"
	"net/http"
)

func main() {
	// This is a placeholder for the main function.
	// You can add your code here to run the application.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./home.html")
	})

	http.ListenAndServe(":8080", nil)
	// This is a simple HTTP server that responds with "Hello, World!".
	// You can test it by visiting http://localhost:8080 in your web browser.

}
