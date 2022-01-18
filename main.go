package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type Response struct {
	ShortURL string `json:"short_url"`
}

func shorten_url() string {
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
			random := shorten_url()
			data := Response{
				ShortURL: random,
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":8080", mux)
}

