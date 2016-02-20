package main

import (
	"fmt"
	"io"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
	fmt.Print("Hello World")
}

func main() {
	http.HandleFunc("/login/", login)
	http.ListenAndServe(":1337", nil)
}
