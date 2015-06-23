package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Listener int

func (l *Listener) Printline(line []byte, ack *bool) error {
	fmt.Println(string(line))
	return nil
}

func main() {
	// ResolveTCPAddr parses addr as a TCP address of the form "host:port"
	// or "[ipv6-host%zone]:port" and resolves a pair of domain name and
	// port name on the network net, which must be "tcp", "tcp4" or "tcp6".
	laddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42585")
	if err != nil {
		log.Fatal(err)
	}

	// ListenTCP announces on the TCP address laddr and returns a TCP
	// listener. Net must be "tcp", "tcp4", or "tcp6". If laddr has a
	// port of 0, ListenTCP will choose an available port
	inbound, err := net.ListenTCP("tcp", laddr)

	if err != nil {
		log.Fatal(err)
	}

	listener := new(Listener)

	// Register publishes the receiver's methods in the DefaultServer.
	rpc.Register(listener)

	// Accept accepts connections on the listener and serves requests to
	// DefaultServer for each incoming connection.
	rpc.Accept(inbound)
}
