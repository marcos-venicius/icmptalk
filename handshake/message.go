package handshake

import (
	"net"
)

func (h *handshake) sendMessage(message string, to string) error {
	packet, err := createIcmpPacket([]byte(message))

	if err != nil {
		return err
	}

	if _, err := h.conn.WriteTo(packet, &net.IPAddr{IP: net.ParseIP(to)}); err != nil {
		return err
	}

	return nil
}

// listenMessage ip, data, error
func (h *handshake) listenMessage() (string, []byte, error) {
	msg := make([]byte, 64)

	length, ip, err := h.conn.ReadFrom(msg)

	msg = msg[:length]

	return ip.String(), msg, err
}
