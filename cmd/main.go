package main

import (
	"fmt"
	"github.com/Komei22/sqd/server"
	"net"
	"os"
)

func main() {
	lAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:1234")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	s := server.New(lAddr)
	s.Start()
}
