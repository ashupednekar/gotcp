package main

import (
	"fmt"

	"github.com/ashupednekar/gotcp/server"
)

func main() {
	s := server.NewServer(":3001")

	go func() {
		for message := range s.Chans.Msgchan {
			fmt.Printf("received message: %s from %s", string(message.Payload), message.Source)
		}
	}()

	s.Start()
}
