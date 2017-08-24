package main

import (
	"flag"
	"fmt"
	"github.com/m4dfry/go-term-utils/utils"
)

func usage() {
	fmt.Printf("Usage: go-term-utils [options] command <params>\n")
	fmt.Printf("Commands & params: \n")
	fmt.Printf("  curr    <convert_from> <convert_to> [<value>]\n")
	fmt.Printf("  crycurr <convert_to>\n")
	fmt.Printf("  ip      [<address>]\n")
	fmt.Printf("  space\n")
	fmt.Printf("  tor\n")
}

func main() {
	var config = flag.String("c", "", "Custom configuration file path")

	flag.Usage = func() {
		usage()
		flag.PrintDefaults()
	}

	flag.Parse()
	cmd := flag.Arg(0)

	if *config != "" {
		fmt.Println("Find custom config on:", &config)
	}

	switch cmd {
	case "curr":
		utils.Currency(flag.Args())
	case "crycurr":
		utils.CryptoCurrency(flag.Args())
	case "ip":
		utils.IpInfo(flag.Args())
	case "space":
		utils.Space()
	case "tor":
		utils.Tor()
	default:
		fmt.Println("unrecognized command")
		usage()
		return
	}
}
