package handshake

import (
	"log"

	"golang.org/x/net/icmp"
)

type handshake struct {
	step    int
	numbers []int
	steps   int
	ip      string
	iface   string
	conn    *icmp.PacketConn
}

func NewHandshake(iface string) *handshake {
	steps := 4

	conn, err := icmp.ListenPacket("ip4:icmp", iface)

	if err != nil {
		log.Fatal(err)
	}

	return &handshake{
		step:    0,
		iface:   iface,
		numbers: make([]int, steps),
		steps:   steps,
		conn:    conn,
		ip:      "",
	}
}

func (h *handshake) addStep(step int, ip string) bool {
	if h.steps == 0 || h.shouldValidate() {
		return false
	}

	if h.step == 0 {
		h.ip = ip
	} else if h.ip != ip {
		return false
	}

	h.numbers[h.step] = step
	h.step++

	return true
}

func (h *handshake) shouldValidate() bool {
	return h.step == h.steps
}

func (h *handshake) sumSteps() int {
	s := 0

	for _, n := range h.numbers[:h.steps-1] {
		s += n
	}

	return s
}

func (h *handshake) validate() bool {
	s := h.sumSteps()

	return s == h.numbers[h.steps-1] && s%2 != 0
}
