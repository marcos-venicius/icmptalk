package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/marcos-venicius/icmptalk/handshake"
)

func main() {
	iface := flag.String("iface", "0.0.0.0", "your machine ip")

	flag.Parse()

	addr, err := handshake.ListenForConnection(*iface)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfull handshake with %s\n", addr)
}
