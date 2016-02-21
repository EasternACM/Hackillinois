package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func getCode(w http.ResponseWriter, r *http.Request) {
	log.Print("hit")
	code := r.URL.Query().Get("code")
	log.Print(code)
	sendCode(code)
}

func sendCode(code string) {
	client := &http.Client{}
	var jsonStr = []byte("grant_type=authorization_code&redirect_uri=http://zacc.xyz:8000&code=" + code)
	req, err := http.NewRequest("POST", "https://hackillinois.climate.com/api/oauth/token", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Basic ZHBxazVzbXBxMDM5Mmo6dDB0czB0YWdvcm05bnExdjZzbW10dnBxYzI=")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "*/*")
	res, err1 := client.Do(req)
	if err1 != nil {
		panic(err)
	}
	defer res.Body.Close()
	//Prints the response
	log.Print(res)
}

func getFields(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	log.Print("start")
	http.HandleFunc("/", getCode)
	http.HandleFunc("/fields", getFields)
	http.ListenAndServe(":8000", nil)
}
