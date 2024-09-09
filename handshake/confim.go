package handshake

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func (h *handshake) confirm() error {
	s := h.sumSteps() * 2

	message := fmt.Sprintf("|%d|", s)

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: []byte(message),
		},
	}

	wb, err := wm.Marshal(nil)

	if err != nil {
		return err
	}

	if _, err := h.conn.WriteTo(wb, &net.IPAddr{IP: net.ParseIP(h.ip)}); err != nil {
		return err
	}

	return nil
}
