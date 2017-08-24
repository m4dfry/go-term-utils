package utils

import (
	"fmt"
)

// SCheckTorprojectOrg struct
type SCheckTorprojectOrg struct {
	IsTor bool   `json:"IsTor"`
	IP    string `json:"IP"`
}

// Tor main function
func Tor() {
	var check SCheckTorprojectOrg
	TorAPICall("https://check.torproject.org/api/ip", &check)

	fmt.Println("Is Tor?  ", check.IsTor)
	fmt.Println("IP     : ", check.IP)
}
