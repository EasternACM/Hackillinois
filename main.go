package main

import (
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	log.Print("hit")
	code := r.URL.Query().Get("code")
	log.Print(code)
}

func main() {
	log.Print("start")
	http.HandleFunc("/", login)
	http.ListenAndServe(":8000", nil)
}
