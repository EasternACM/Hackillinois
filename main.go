package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	fmt.Print(code)
}

func main() {
	http.HandleFunc("/login", login)
	http.ListenAndServe(":1337", nil)
}
