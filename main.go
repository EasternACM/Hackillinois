package main

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Print("logged")
}

func main() {
	http.HandleFunc("/login/", login)
	http.ListenAndServe(":8000", nil)
}
