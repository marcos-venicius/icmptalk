package handshake

import (
	"errors"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/icmp"
)

// ListenForConnection returns (connection net address, error)
func ListenForConnection(iface string) (net.Addr, error) {
	conn, err := icmp.ListenPacket("ip4:icmp", iface)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Printf("[*] listening for connections at %s...\n", iface)

	hs := newHandshake()

	msg := make([]byte, 64)

	for {
		length, sourceIP, err := conn.ReadFrom(msg)

		if err != nil {
			log.Fatal(err)
		}

		n, err := parseGreeting(msg[:length])

		if err != nil {
			fmt.Printf("[%s] Invalid greeting\n", sourceIP)
			hs = newHandshake()
			continue
		}

		if hs.addStep(n, sourceIP.String()) {
			fmt.Printf("[%s] (%d/%d)\n", sourceIP, hs.step, hs.steps)
		} else {
			fmt.Printf("[%s] Invalid handshake\n", sourceIP)
			hs = newHandshake()
			continue
		}

		if hs.shouldValidate() {
			fmt.Printf("[%s] Validating handshake\n", sourceIP)

			isValid := hs.validate()

			if isValid {
				return sourceIP, nil
			}

			break
		}
	}

	return nil, errors.New("no valid handshake")
}
