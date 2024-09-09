package handshake

import (
	"log"

	"golang.org/x/net/icmp"
)

type Handshake struct {
	step    int
	numbers []int
	steps   int
	ip      string
	iface   string
	conn    *icmp.PacketConn
}

func NewHandshake(iface string) *Handshake {
	steps := 4

	conn, err := icmp.ListenPacket("ip4:icmp", iface)

	if err != nil {
		log.Fatal(err)
	}

	return &Handshake{
		step:    0,
		iface:   iface,
		numbers: make([]int, steps),
		steps:   steps,
		conn:    conn,
		ip:      "",
	}
}

func (h *Handshake) Close() {
	h.conn.Close()
}

func (h *Handshake) Connection() *icmp.PacketConn {
	return h.conn
}

func (h *Handshake) Pair() string {
	return h.ip
}

func (h *Handshake) addStep(step int, ip string) bool {
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

func (h *Handshake) shouldValidate() bool {
	return h.step == h.steps
}

func (h *Handshake) sumSteps() int {
	s := 0

	for _, n := range h.numbers[:h.steps-1] {
		s += n
	}

	return s
}

func (h *Handshake) validate() bool {
	s := h.sumSteps()

	return s == h.numbers[h.steps-1] && s%2 != 0
}
