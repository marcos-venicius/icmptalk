package handshake

import (
	"os"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// createIcmpPacket return the icmp packet in bytes
func createIcmpPacket(data []byte) ([]byte, error) {
	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  1,
			Data: data,
		},
	}

	wb, err := wm.Marshal(nil)

	if err != nil {
		return nil, err
	}

	return wb, nil
}
