package main

import (
	_ "errors"
	"fmt"
	"io"
	"net/http"
	_ "os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	_, err := io.WriteString(w, "This is my website!\n")
	if err != nil {
		return
	}
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	_, err := io.WriteString(w, "Hello, HTTP!\n")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	_ = http.ListenAndServe(":8080", nil)
}
