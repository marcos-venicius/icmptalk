package handshake

import "strconv"

func parseHandshakeStep(msg []byte) (int, error) {
	d := parseString(msg)

	n := string(d)

	return strconv.Atoi(n)
}

func parseString(msg []byte) string {
	s, c, d := false, 0, make([]byte, 64)

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

	return string(d[:c])
}
