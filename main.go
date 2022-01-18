package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

//json format struct
type Response struct {
	ShortURL string `json:"short_url"`
}

//return shortened url
func shorten_url() string {

	//create random string of length 8
	bytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	//new shortened url
	var random string
	random = "https://127.0.0.1/" + (string(bytes))
	return random
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/secureworks.com", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//get ranomized short url
			random := shorten_url()

			//put random url in json format
			data := Response{
				ShortURL: random,
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	//open server on port 8080
	http.ListenAndServe(":8080", mux)
}
