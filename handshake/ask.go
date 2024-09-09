package handshake

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net"

	"github.com/marcos-venicius/icmptalk/protocol"
)

func (h *Handshake) AskForConnection(ip string) error {
	parsed := net.ParseIP(ip)

	if parsed == nil {
		return errors.New("invalid ip address")
	}

	h.ip = ip

	i, s, numbers := 0, 0, make([]int, h.steps)

	for i < h.steps-2 {
		n := rand.IntN(100)

		numbers[i] = n

		s += n
		i++
	}

	if s%2 == 0 {
		numbers[i] = 1
	}

	numbers[i+1] = s + numbers[i]

	for _, n := range numbers {
		message := fmt.Sprintf("|%d|", n)

		err := protocol.SendMessage(h.conn, message, h.ip)

		if err != nil {
			return err
		}

		sourceIP, _, err := protocol.ListenMessage(h.conn, 64)

		if err != nil {
			return err
		}

		if sourceIP != ip {
			return errors.New("invalid responder")
		}
	}

	fmt.Println("handshake sent, waiting confirmation...")

	sourceIP, response, err := protocol.ListenMessage(h.conn, 64)

	if err != nil {
		return err
	}

	if sourceIP != ip {
		return errors.New("uknown pair response")
	}

	n, err := parseHandshakeStep(response)

	if err != nil {
		return err
	}

	expected := numbers[len(numbers)-1] * 2

	if n != expected {
		protocol.SendMessage(h.conn, "|FAIL|", h.ip)

		return errors.New("invalid handshake")
	}

	protocol.SendMessage(h.conn, "|OK|", h.ip)

	return nil
}
