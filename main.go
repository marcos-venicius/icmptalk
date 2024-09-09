package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/marcos-venicius/icmptalk/chat"
	"github.com/marcos-venicius/icmptalk/handshake"
)

func main() {
	iface := flag.String("iface", "0.0.0.0", "your machine ip")
	listenMode := flag.Bool("listen", false, "starts in listen mode")
	target := flag.String("target", "", "your target")

	flag.Parse()

	hs := handshake.NewHandshake(*iface)

	if *listenMode {
		err := hs.ListenForConnection()

		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := hs.AskForConnection(*target)

		if err != nil {
			log.Fatal(err)
		}
	}

	defer hs.Close()

	fmt.Println("Connected successfully")

	c := chat.NewChat(hs)

	fmt.Println("Chat started")

	go c.Listen()

	go func() {
		for {
			message := <-c.Messages()

			if message.Me {
				fmt.Printf("\033[1;33m# \033[0m %s\n", message.Content)
			} else {
				fmt.Printf("\033[1;33m# \033[0m %s\n", message.Content)
			}
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)

		text, _ := reader.ReadString('\n')

		err := c.Send(text)

		if err != nil {
			fmt.Printf("\033[1;31mE \033[0m %s\n", err)
		}
	}
}
