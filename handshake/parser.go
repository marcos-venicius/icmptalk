package handshake

import "strconv"

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
