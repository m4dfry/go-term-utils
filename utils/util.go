package utils

import (
	"encoding/json"
	"golang.org/x/net/proxy"
	"log"
	"net/http"
	"net/url"
)

// SafeParam function
func SafeParam(s string) string {
	return url.QueryEscape(s)
}

// APICall function
func APICall(apiAddress string, data interface{}) {
	GenericAPICall(apiAddress, data, nil)
}

// TorAPICall function
func TorAPICall(apiAddress string, data interface{}) {
	// Create a transport that uses Tor Browser's SocksPort.  If
	// talking to a system tor, this may be an AF_UNIX socket, or
	// 127.0.0.1:9050 instead.
	tbProxyURL, err := url.Parse("socks5://127.0.0.1:9050")
	if err != nil {
		log.Fatal("Failed to parse proxy URL: %v\n", err)
	}

	// Get a proxy Dialer that will create the connection on our
	// behalf via the SOCKS5 proxy.  Specify the authentication
	// and re-create the dialer/transport/client if tor's
	// IsolateSOCKSAuth is needed.
	tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
	if err != nil {
		log.Fatal("Failed to obtain proxy dialer: %v\n", err)
	}

	// Make a http.Transport that uses the proxy dialer, and a
	// http.Client that uses the transport.
	tbTransport := &http.Transport{Dial: tbDialer.Dial}

	GenericAPICall(apiAddress, data, tbTransport)
}

// GenericAPICall function
func GenericAPICall(address string, data interface{}, tbTransport *http.Transport) {

	log.Println("urlxs:", address)
	// Build the request
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}
	if tbTransport != nil {
		client = &http.Client{Transport: tbTransport}
	}

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
