package handshake

import (
	"errors"
	"fmt"
	"golang.org/x/net/icmp"
	"log"
	"net"
)

// ListenForConnection returns (connection net address, error)
func ListenForConnection(iface string) (net.Addr, error) {
	conn, err := icmp.ListenPacket("ip4:icmp", iface)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Printf("[*] listening for connections at %s...\n", iface)

	handshake := NewHandshake()

	msg := make([]byte, 64)

	for {
		length, sourceIP, err := conn.ReadFrom(msg)

		if err != nil {
			log.Fatal(err)
		}

		n, err := parseGreeting(msg[:length])

		if err != nil {
			fmt.Printf("[%s] Invalid greeting\n", sourceIP)
			handshake = NewHandshake()
			continue
		}

		if handshake.addStep(n, sourceIP.String()) {
			fmt.Printf("[%s] (%d/%d)\n", sourceIP, handshake.step, handshake.steps)
		} else {
			fmt.Printf("[%s] Invalid handshake\n", sourceIP)
			handshake = NewHandshake()
			continue
		}

		if handshake.shouldValidate() {
			fmt.Printf("[%s] Validating handshake\n", sourceIP)

			isValid := handshake.validate()

			if isValid {
				return sourceIP, nil
			}

			break
		}
	}

	return nil, errors.New("no valid handshake")
}
