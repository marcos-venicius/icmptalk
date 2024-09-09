package protocol

import (
	"os"

	"net"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

// createIcmpPacket return the icmp packet in bytes
func CreateIcmpPacket(data []byte) ([]byte, error) {
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

func SendMessage(conn *icmp.PacketConn, message, to string) error {
	packet, err := CreateIcmpPacket([]byte(message))

	if err != nil {
		return err
	}

	if _, err := conn.WriteTo(packet, &net.IPAddr{IP: net.ParseIP(to)}); err != nil {
		return err
	}

	return nil
}

// listenMessage ip, data, error
func ListenMessage(conn *icmp.PacketConn, bufsize int) (string, []byte, error) {
	msg := make([]byte, bufsize)

	length, ip, err := conn.ReadFrom(msg)

	msg = msg[:length]

	return ip.String(), msg, err
}
