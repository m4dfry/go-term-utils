package utils

import (
	"fmt"
	"log"
)

// SOpenGeneric struct
type SOpenGeneric struct {
	Message string `json:"message"`
}

// SOpenIssnow source : http://api.open-notify.org/iss-now.json
type SOpenIssnow struct {
	SOpenGeneric
	Timestamp   int `json:"timestamp"`
	IssPosition struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"iss_position"`
}

// SOpenAstros source : http://api.open-notify.org/astros.json
type SOpenAstros struct {
	SOpenGeneric
	Number int `json:"number"`
	People []struct {
		Name  string `json:"name"`
		Craft string `json:"craft"`
	} `json:"people"`
}

// SGoogleGEO source : http://maps.google.com/maps/api/geocode/json?latlng=4.1108,34.5748
type SGoogleGEO struct {
	Results []struct {
		FormattedAddress string `json:"formatted_address"`
	} `json:"results"`
	Status string `json:"status"`
}

// CheckMessage func
func (s *SOpenGeneric) CheckMessage() {
	if s.Message != "success" {
		log.Fatal("ERROR : Fail to load information")
		return
	}
}

// GetLatLong func
func (s *SOpenIssnow) GetLatLong() string {
	return s.IssPosition.Latitude + "," + s.IssPosition.Longitude
}

// Space main function
func Space() {
	var issnow SOpenIssnow
	APICall("http://api.open-notify.org/iss-now.json", &issnow)
	issnow.CheckMessage()

	var geoloc SGoogleGEO
	APICall("http://maps.google.com/maps/api/geocode/json?latlng="+SafeParam(issnow.GetLatLong()), &geoloc)
	//fmt.Printf("--- Google Response\n%+v\n--- END", geoloc)

	var astros SOpenAstros
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
