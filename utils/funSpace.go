package utils

import (
  "fmt"
  "log"
)

type S_open_generic struct {
	Message string `json:"message"`
}

// http://api.open-notify.org/iss-now.json
type S_open_issnow struct {
	S_open_generic
	Timestamp   int `json:"timestamp"`
	IssPosition struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"iss_position"`
}

//http://api.open-notify.org/astros.json
type S_open_astros struct {
	S_open_generic
	Number int `json:"number"`
	People []struct {
		Name  string `json:"name"`
		Craft string `json:"craft"`
	} `json:"people"`
}

//http://maps.google.com/maps/api/geocode/json?latlng=4.1108,34.5748
type S_GoogleGEO struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
	} `json:"results"`
	Status string `json:"status"`
}

func (s *S_open_generic) CheckMessage() {
	if s.Message != "success" {
		log.Fatal("ERROR : Fail to load information")
		return
	}
}

func (s *S_open_issnow) GetLatLong() string {
	return s.IssPosition.Latitude + "," + s.IssPosition.Longitude
}

func Space() {
	var issnow S_open_issnow
	APICall("http://api.open-notify.org/iss-now.json", &issnow)
	issnow.CheckMessage()

	var geoloc S_GoogleGEO
	APICall("http://maps.google.com/maps/api/geocode/json?latlng=" + SafeParam(issnow.GetLatLong()), &geoloc)
	//fmt.Printf("--- Google Response\n%+v\n--- END", geoloc)

	var astros S_open_astros
	APICall("http://api.open-notify.org/astros.json", &astros)
	astros.CheckMessage()

	address := " Unable to load GeoPosition (" + geoloc.Status + ")"
	if len(geoloc.Results) > 0 {
		address = "-> " + geoloc.Results[0].FormattedAddress
	}

	fmt.Println("ISS Position Now:")
	fmt.Printf("Lat(%v) Long(%v)\n%s\n\n", issnow.IssPosition.Latitude, issnow.IssPosition.Longitude, address)

	fmt.Printf("There are %d people in space now:\n", astros.Number)
	for _, astro := range astros.People {
		fmt.Printf(" - %s on %s\n", astro.Name, astro.Craft)
	}
}

