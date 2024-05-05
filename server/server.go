package server

import (
	"log"
	"net"
)

type Server struct {
	ListenAddr string
	ln         net.Listener
	Msgchan    chan Message
	quitchan   chan struct{}
}

type Message struct {
	Source  string
	Payload []byte
}

func NewServer(addr string) *Server {
	return &Server{
		ListenAddr: addr,
		Msgchan:    make(chan Message, 10),
		quitchan:   make(chan struct{}),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln
	s.AcceptLoop()

	<-s.quitchan
	close(s.Msgchan)

	return nil
}

func (s *Server) AcceptLoop() {
	println("Accepting connections at ", s.ListenAddr)
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Fatal("error while accepting: ", err)
		}
		go s.HandleConn(conn)
	}
}
