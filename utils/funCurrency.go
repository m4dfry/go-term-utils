package utils

import (
  "fmt"
  "strconv"
  "log"
)

type S_fixer_io struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func Currency(args []string) {
	if len(args) < 3 {
		fmt.Printf("Error: not enough argument for curr\n\n")
	}
	url := "http://api.fixer.io/latest?base=" + SafeParam(args[1]) + "&symbols=" + SafeParam(args[2])
	var data S_fixer_io
	APICall(url, &data)

	fmt.Println("Conversion :", data.Base, " -> ", args[2])
	fmt.Println("Rate       :", data.Rates[args[2]])

	if len(args) > 3 {
		convs, err := strconv.ParseFloat(args[3], 64)
		if err != nil {
			log.Fatal("ParseFloat: ", err)
			return
		}
		convv := convs * data.Rates[args[2]]
		fmt.Printf("Value      : %v -> %.2f\n", convs, convv)
	}
}

