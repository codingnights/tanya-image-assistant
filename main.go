package main

import (
	"net/http"
	"io/ioutil"
	"bytes"

	"encoding/json"
	"github.com/tidwall/gjson"
)

type Message struct {
	Requests []Request`json:"requests"`
}

type Request struct {
	Images Image`json:"images"`
	Features []Feature`json:"features"`
}

type Image struct {
	Content string`json:"content"`
}

type Feature struct{
	Type string `json:"type"`
}


func recognize(w http.ResponseWriter, r *http.Request){
	var key = ""
	var url = "https://vision.googleapis.com/v1/images:annotate?key=" + key


	body, _ := ioutil.ReadAll(r.Body)
	var name = gjson.Get(body, "img")
	var feature = Feature{"TEXT_DETECTION"}
	var image = Image{name}
	var request = Request{image, {feature}}
	var message = Message{{request}}

	var jsonStr, erj= json.Marshal(message)
	if (erj != nil){w.WriteHeader(500)}


	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if(err != nil){w.WriteHeader(500)}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, er2 := client.Do(req)
	if er2 != nil {
		w.WriteHeader(200)
	}
	defer resp.Body.Close()


	w.WriteHeader(200)

}

func main (){
	http.HandleFunc("/cha-ching", recognize)
}