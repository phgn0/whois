// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/domainr/go-whois/whois"
	_ "github.com/domainr/go-whois/servers"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [arguments] <domain>\n\nAvailable arguments:\n", os.Args[0])
		flag.PrintDefaults()

		os.Exit(1)
	}

	flag.Parse()
	query := flag.Arg(0)

	if query == "" {
		flag.Usage()
	}

	req, err := whois.Resolve(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := req.Fetch()
	if err != nil {
		fmt.Println(err)
		return
	}

	res.WriteMIME(os.Stdout)
}
