package handshake

import (
	"errors"
	"fmt"
	"log"
	"net"
)

// ListenForConnection returns (connection net address, error)
func (h *handshake) ListenForConnection() (net.Addr, error) {
	defer h.conn.Close()

	fmt.Printf("[*] listening for connections at %s...\n", h.iface)

	msg := make([]byte, 64)

	for {
		length, sourceIP, err := h.conn.ReadFrom(msg)

		if err != nil {
			log.Fatal(err)
		}

		n, err := parseHandshakeStep(msg[:length])

		if err != nil {
			fmt.Printf("[%s] Invalid greeting\n", sourceIP)
			break
		}

		if h.addStep(n, sourceIP.String()) {
			fmt.Printf("[%s] (%d/%d)\n", sourceIP, h.step, h.steps)
		} else {
			fmt.Printf("[%s] Invalid handshake\n", sourceIP)
			break
		}

		if h.shouldValidate() {
			fmt.Printf("[%s] Validating handshake\n", sourceIP)

			isValid := h.validate()

			if !isValid {
				break
			}

			s := h.sumSteps() * 2

			err := h.sendMessage(fmt.Sprintf("|%d|", s), h.ip)

			if err != nil {
				return nil, err
			}

			return sourceIP, nil
		}
	}

	return nil, errors.New("no valid handshake")
}
