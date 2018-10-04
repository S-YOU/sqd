package server

import (
	"fmt"
	"net"
)

type Server struct {
	listenAddr *net.TCPAddr
}

// New return Server object
func New(lAddr *net.TCPAddr) *Server {
	return &Server{listenAddr: lAddr}
}

// Start run server
func (s *Server) Start() error {
	listener, err := net.ListenTCP("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Print("client message:")
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			return err
		}
		fmt.Print(string(buf[:n]))
	}
	return nil
}
