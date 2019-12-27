package telnet

import (
	"log"
	"net"
)

type Handler interface {
	ServeTELNET(ts *TelnetStream)
}

type HandlerFunc func() (int, HandlerFunc, error)

type Server struct {
	Addr    string
	Handler Handler
}

func (server *Server) ListenAndServe() error {

	// initialze a default handler
	if server.Handler == nil {
		server.Handler = NewEchoHandler()
	}

	var clients []net.Conn

	taddr, err := net.ResolveTCPAddr("tcp", server.Addr)
	if err != nil {
		log.Println(err)
		return err
	}

	listener, err := net.ListenTCP("tcp", taddr)
	if err != nil {
		log.Println(err)
		return err
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue // keep going despite this one error
		}

		ip, err := net.ResolveTCPAddr("tcp", conn.LocalAddr().String())
		log.Println("Connection from", ip)

		clients = append(clients, conn)

		ts := NewTelnetStream(conn)
		go server.Handler.ServeTELNET(ts)
	}

	return nil
}

func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
