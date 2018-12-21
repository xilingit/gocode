package main

import (
	"../add"
	"fmt"
	"net/http"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path:", r.URL.Path)
	fmt.Fprintf(w, add.Name)
}

func main() {

	var a int = 9223372036854775807
	fmt.Println(a)
	var b byte = 'c'
	fmt.Println(b)
}
