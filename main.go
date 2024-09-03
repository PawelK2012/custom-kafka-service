package main

import (
	"fmt"
	"io"
	"log/slog"
	"net"
)

type Header struct {
	Size int32
}
type Message struct {
	data []byte
}

type Server struct {
	// coffsets map[string]int
	// buffer   []Message
	ln net.Listener
}

func NewServer() *Server {
	return &Server{}

}

func (s *Server) Start() error {
	return nil
}

func (s *Server) Listen() error {
	ln, err := net.Listen("tcp", ":9092")
	if err != nil {
		return err
	}
	s.ln = ln
	for {
		conn, err := ln.Accept()
		if err != nil {
			if err == io.EOF {
				return err
			}
			slog.Error("server accept error", "err", err)
		}
		go s.hanldeConn(conn)
	}
}

func (s *Server) hanldeConn(conn net.Conn) {
	fmt.Println("new connection", conn.RemoteAddr())
}

func main() {
	fmt.Println("Hello, 世界")
	server := NewServer()
	server.Listen()
}
