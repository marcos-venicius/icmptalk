package handshake

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net"
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

		err := h.sendMessage(message)

		if err != nil {
			return err
		}

		sourceIP, _, err := h.listenMessage()

		if err != nil {
			return err
		}

		if sourceIP != ip {
			return errors.New("invalid responder")
		}
	}

	fmt.Println("handshake sent, waiting confirmation...")

	sourceIP, response, err := h.listenMessage()

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
		h.sendMessage("|FAIL|")

		return errors.New("invalid handshake")
	}

	h.sendMessage("|OK|")

	return nil
}
