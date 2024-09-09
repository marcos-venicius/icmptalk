package chat

import (
	"fmt"

	"github.com/marcos-venicius/icmptalk/handshake"
	"github.com/marcos-venicius/icmptalk/protocol"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type Message struct {
	Me      bool
	Content string
}

type Chat struct {
	h        *handshake.Handshake
	messages chan Message
}

func NewChat(h *handshake.Handshake) *Chat {
	return &Chat{
		h:        h,
		messages: make(chan Message),
	}
}

func (c *Chat) Send(message string) error {
	err := protocol.SendMessage(c.h.Connection(), message, c.h.Pair())

	protocol.ListenMessage(c.h.Connection(), 1)

	return err
}

func (c *Chat) Listen() {
	conn := c.h.Connection()

	for {
		buf := make([]byte, 1024)

		_, ip, err := conn.ReadFrom(buf)

		if err != nil {
			fmt.Printf("\033[0;31mERROR\033[0m %s\n", err)
			continue
		}

		if ip.String() != c.h.Pair() {
			continue // ignore who did not did the handshake
		}

		packet, err := icmp.ParseMessage(1, buf)

		if err != nil {
			fmt.Printf("\033[0;31mERROR\033[0m %s\n", err)
			continue
		}

		if packet.Code == 0 && packet.Type == ipv4.ICMPTypeEcho {
			message := Message{
				Me:      ip.String() != c.h.Pair(),
				Content: string(packet.Body.(*icmp.Echo).Data),
			}

			c.messages <- message
		}

	}
}

func (c *Chat) Messages() <-chan Message {
	return c.messages
}
