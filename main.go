package main

import (
	"bytes"
	"encoding/json"
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
	log.Println(req)
	res, err1 := client.Do(req)
	if err1 != nil {
		panic(err)
	}
	defer res.Body.Close()
	log.Println(res)
}

func getFields(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	myItems := []string{"item1", "item2", "item3"}
	a, _ := json.Marshal(myItems)

	w.Write(a)

}

func main() {
	log.Print("start")
	http.HandleFunc("/", sendCode)
	http.HandleFunc("/fields", getFields)
	http.ListenAndServe(":8000", nil)
}
