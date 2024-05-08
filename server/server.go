package server

import (
	"log"
	"net"
)

type Channels struct {
	Msgchan  chan Message
	quitchan chan struct{}
}

type Server struct {
	ListenAddr string
	ln         net.Listener
	Chans      Channels
}

type Message struct {
	Source  string
	Payload []byte
}

func NewServer(addr string) *Server {
	return &Server{
		ListenAddr: addr,
		Chans: Channels{
			Msgchan:  make(chan Message, 10),
			quitchan: make(chan struct{}),
		},
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln
	s.AcceptLoop()

	<-s.Chans.quitchan
	close(s.Chans.Msgchan)

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
