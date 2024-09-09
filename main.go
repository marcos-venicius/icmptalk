package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/marcos-venicius/icmptalk/handshake"
)

func main() {
	iface := flag.String("iface", "0.0.0.0", "your machine ip")
	listenMode := flag.Bool("listen", false, "starts in listen mode")
	target := flag.String("target", "", "your target")

	flag.Parse()

	hs := handshake.NewHandshake(*iface)

	if *listenMode {
		addr, err := hs.ListenForConnection()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Successfull handshake with %s\n", addr)
	} else {
		err := hs.AskForConnection(*target)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Connected successfully")
	}
}
