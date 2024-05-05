package server

import (
	"fmt"
	"log"
	"net"
)

func (s *Server) HandleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal("error handling conn: ", err)
			continue
		}
		msg := buf[:n]
		println("msg: ", string(msg))
		fmt.Fprintf(conn, "Thanks, received %s", string(msg))
	}
}
