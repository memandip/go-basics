package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>Hello GoLang!!!</h1>")
}

func about(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>About Page</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Go server listening on port: 3000")
	http.ListenAndServe(":3000", nil)
}
