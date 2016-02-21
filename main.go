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
	log.Println(res.Body)
}

func getFields(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	myItems := `{
	    name: Back 40,
	    source: {
	        attributes: FIELDVIEW,
	        boundary: EQUIPMENT_GENERATED
	    },
	    boundary: null,
	    sourceMetadata: meta,
	    farmId: 0ec15ec6-3c28-472f-8d42-012d3e93221c,
	    centroid: {
	    type: Point,
	        coordinates: [
	          -101.82513480004555,
	          33.60125723889728
	        ]
	    },
	    id: 24578672,
	    area: {
	        u: ac,
	        q: 128.51
	    },
	    uuid: 5d488989-7898-43c1-92f4-763a2dfe728c,
	    version: 1,
	    timestamp: 1453486212000,
	    enabled: true
	}`
	a, _ := json.Marshal(myItems)
	w.Write(a)

}

func main() {
	log.Print("start")
	http.HandleFunc("/", sendCode)
	http.HandleFunc("/fields", getFields)
	http.ListenAndServe(":8000", nil)
}
