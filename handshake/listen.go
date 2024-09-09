package handshake

import (
	"errors"
	"fmt"
	"log"
	"net"
)

// ListenForConnection returns (connection net address, error)
func ListenForConnection(iface string) (net.Addr, error) {
	hs := newHandshake(iface)

	defer hs.conn.Close()

	fmt.Printf("[*] listening for connections at %s...\n", iface)

	msg := make([]byte, 64)

	for {
		length, sourceIP, err := hs.conn.ReadFrom(msg)

		if err != nil {
			log.Fatal(err)
		}

		n, err := parseGreeting(msg[:length])

		if err != nil {
			fmt.Printf("[%s] Invalid greeting\n", sourceIP)
			break
		}

		if hs.addStep(n, sourceIP.String()) {
			fmt.Printf("[%s] (%d/%d)\n", sourceIP, hs.step, hs.steps)
		} else {
			fmt.Printf("[%s] Invalid handshake\n", sourceIP)
			break
		}

		if hs.shouldValidate() {
			fmt.Printf("[%s] Validating handshake\n", sourceIP)

			isValid := hs.validate()

			if !isValid {
				break
			}

			err := hs.confirm()

			if err != nil {
				return nil, err
			}

			return sourceIP, nil
		}
	}

	return nil, errors.New("no valid handshake")
}
