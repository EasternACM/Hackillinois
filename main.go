package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type Response struct {
	XCache        string
	XAmzCfId      string
	ContentType   string
	ContentLength int
	Connection    string
	Date          string
}

func sendCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
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
	log.Print(res)
	var respo Response
	//Prints the response
	log.Println("\n\n XamzCfId" + respo.XAmzCfId)
	log.Println("Xcache\n" + respo.XCache)
	log.Println("\nDate\n" + respo.Date)
}

func getFields(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	log.Print("start")
	http.HandleFunc("/", sendCode)
	http.HandleFunc("/fields", getFields)
	http.ListenAndServe(":8000", nil)
}
