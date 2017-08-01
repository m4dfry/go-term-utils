package utils

import (
  "fmt"
)

type S_ip_api_com struct {
	As          string  `json:"as"`
	City        string  `json:"city"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	ISP         string  `json:"isp"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Org         string  `json:"org"`
	Query       string  `json:"query"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	Status      string  `json:"status"`
	Timezone    string  `json:"timezone"`
	Zip         string  `json:"zip"`
}

func IpInfo(args []string) {
	argIp := ""
	if len(args) > 1 {
		argIp = args[1]
	}

	var info S_ip_api_com
	APICall("http://ip-api.com/json/"+argIp, &info)

	fmt.Println("IP       :", info.Query)
	fmt.Println("ISP      :", info.ISP)
	fmt.Printf("Location : %s(%s), %s(%s)\n", info.City, info.Zip, info.Country, info.CountryCode)

}
