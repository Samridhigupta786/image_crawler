package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type image_urls struct {
	Urls []string
}

func submitRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var payload image_urls
	fmt.Printf("body is %v\n", body)
	err = json.Unmarshal(body, &payload)

	res := searchForImageLinks(payload.Urls)
	json.NewEncoder(w).Encode(res)
}
