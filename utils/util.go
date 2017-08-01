package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func SafeParam(s string) string {
	return url.QueryEscape(s)
}

func APICall(url string, data interface{}) {

	log.Println("urlxs:", url)
	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		log.Println(err)
	}
}

