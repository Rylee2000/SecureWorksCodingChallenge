package main

import (
	"net/http"
	"encoding/json"
)

//given a long url, return a shortened version of it with randomly generated ID
func shorten_url(string url) {
	//create the random string ID of length 8
	bytes := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	//new shortened url
	var random string;
	random := "https://127.0.0.1/" + (string(bytes))

	//Put the new url in json format
	data := structs.response{
		ShortURL: random
	}

	return data
}


func main() {
	var long_url string
	var json_data string
	
	//map will store short url to long url combos
	url_map := make(map[string]string)

	//take user input to get long url
	long_url := fmt.Scanln($long_url)
	mux := http.NewServeMux()

	//when user types "/shorten" in url a shortened version of longurl will be shown to them
	mux.handleFunc("/shorten"), func(w http.ResponseWriter, r *http.Request) {
		data := shorten_url(long_url)

		//create a json response
		json_data = structs.Response{
			ShortURL: data
		}
		json.NewEncoder(w).Encode(json_data)
	)}

	//add new shortened url to map
	url_map[data] = long_url

	//when user types shortened url, they will be redirected to the same page as the long url
	mux.handleFunc(json_data[ShortURL]), func(w http.ResponseWriter, r *http.Request) {
		if len(json_data[ShortURL]) > 0 {
			http:Redirect(w, r, url_map[json_data[ShortURL]], http.StatusFound)
		}

	)}

	//open port 8080
	http.ListenAndServe("localhost:8080", mux)
}
