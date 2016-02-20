package main

import (
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
	resp, err := http.Post("https://hackillinois.climate.com/api/oauth/token", "grant_type=\"authorization_code\" scope=\"openid user\" redirect_uri = zacc.xyz:8000 code = "+code, nil)
	if err != nil {
		panic(err)
	}
	log.Print(resp)
}

func main() {
	log.Print("start")
	http.HandleFunc("/", getCode)
	http.ListenAndServe(":8000", nil)

}
