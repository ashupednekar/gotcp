package main

import "github.com/ashupednekar/gotcp/server"

func main() {
	s := server.NewServer(":3001")
	s.Start()
}
