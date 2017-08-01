package main

import (
	"flag"
	"fmt"
	"github.com/m4dfry/go-term-utils/utils"
)

func main() {
	var config = flag.String("c", "", "Custom configuration file path")

	flag.Usage = func() {
		fmt.Printf("Usage: futil [options] command <params>\n\n")
		fmt.Printf("Commands & params: \n\n")
		fmt.Printf("  curr <convert_from> <convert_to> [<value>]\n\n")
		fmt.Printf("  crycurr <convert_to>\n\n")
		fmt.Printf("  ip [<address>]\n\n")
		fmt.Printf("  space\n\n")
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
	default:
		panic("unrecognized command")
	}
}

