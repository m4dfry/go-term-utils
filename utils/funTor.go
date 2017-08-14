package utils

import (
  "fmt"
)
type S_check_torproject_org struct {
  IsTor bool   `json:"IsTor"`
  IP    string `json:"IP"`
}

func Tor() {
	var check S_check_torproject_org
  TorAPICall("https://check.torproject.org/api/ip", &check)

	fmt.Println("Is Tor?  ", check.IsTor)
	fmt.Println("IP     : ", check.IP)
}

