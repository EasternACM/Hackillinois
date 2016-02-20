package main

import (
	"log"
	"net/http"
)

func getCode(w http.ResponseWriter, r *http.Request) {
	log.Print("hit")
	code := r.URL.Query().Get("code")
	log.Print(code)
	sendCode(code, w)
}

func sendCode(code string, w http.ResponseWriter) {
	resp, err := http.Post("https://hackillinois.climate.com/api/oauth/token", "grant_type=authorization_code&redirect_uri=http://zacc.xyz:8000&code="+code, nil)
	if err != nil {
		panic(err)
	}
	log.Print(resp)
	w.Header().Set("Authorization", "Basic ZHBxazVzbXBxMDM5Mmo6dDB0czB0YWdvcm05bnExdjZzbW10dnBxYzI=")
}

func main() {
	log.Print("start")
	http.HandleFunc("/", getCode)
	http.ListenAndServe(":8000", nil)
}
