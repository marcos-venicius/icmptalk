package handshake

import "strconv"

type Handshake struct {
	step    int
	numbers []int
	steps   int
	ip      string
}

func NewHandshake() *Handshake {
	steps := 4

	return &Handshake{
		step:    0,
		numbers: make([]int, steps),
		steps:   steps,
		ip:      "",
	}
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

func (h *Handshake) validate() bool {
	s := 0

	for _, n := range h.numbers[:h.steps-1] {
		s += n
	}

	return s == h.numbers[h.steps-1] && s%2 != 0
}

func parseGreeting(msg []byte) (int, error) {
	s, c, d := false, 0, make([]byte, 5, 5)

	for _, x := range msg {
		if x == '|' && !s {
			s = true
			continue
		} else if x == '|' && s && c == 0 {
			continue
		} else if x == '|' && s {
			break
		} else if s {
			d[c] = x
			c++
		}
	}

	n := string(d[:c])

	return strconv.Atoi(n)
}
