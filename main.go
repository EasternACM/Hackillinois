package main

import (
	"bytes"
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
	client := &http.Client{}
	var jsonStr = []byte("grant_type=authorization_code&redirect_uri=http://zacc.xyz:8000&code=" + code)
	req, err := http.NewRequest("POST", "https://hackillinois.climate.com/api/oauth/token", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Basic ZHBxazVzbXBxMDM5Mmo6dDB0czB0YWdvcm05bnExdjZzbW10dnBxYzI=")
	res, err1 := client.Do(req)
	if err1 != nil {
		panic(err)
	}
	log.Print(res)
}

func main() {
	log.Print("start")
	http.HandleFunc("/", getCode)
	http.ListenAndServe(":8000", nil)
}
