package main

import (
	"flag"
	"log"

	"go.x2ox.com/whois"
)

var domain = flag.String("d", "ip.x2ox.com", "domain name")

func main() {
	flag.Parse()

	d, err := whois.Parse(*domain)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(d.Query())
}
