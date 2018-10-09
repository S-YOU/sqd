package server

import (
	"fmt"
	"net"
	// "github.com/k1LoW/tcpdp/dumper"
	// "github.com/k1LoW/tcpdp/dumper/mysql"
)

// Server struct
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

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			return err
		}
		defer conn.Close()

		go handleConn(conn)
	}
}

func handleConn(conn *net.TCPConn) error {
	defer conn.Close()
	// d := mysql.NewDumper()
	// direction := dumper.ClientToRemote
	// connMetadata := d.NewConnMetadata()
	// additional := []dumper.DumpValue{}

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
		b := buf[:n]
		fmt.Print(string(b))
		// d.Dump(b, direction, connMetadata, additional)
	}

	return nil
}
