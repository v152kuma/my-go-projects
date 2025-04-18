package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/menu", Handler)
	http.ListenAndServe(":4000", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open("./menu.txt")
	io.Copy(w, f)
}
