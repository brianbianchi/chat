package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type msg struct {
	msg  string
	conn net.Conn
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("Unable to start server: %s", err.Error())
	}

	// Create channels for incoming connections, dead connections, and messages
	aconns := make(map[net.Conn]int)
	conns := make(chan net.Conn)
	dconns := make(chan net.Conn)
	msgs := make(chan msg)
	n := 1

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Println(err.Error())
			}
			conns <- conn
		}
	}()
	for {
		select {
		case conn := <-conns:
			aconns[conn] = n
			n++
			// Read messages from connection
			go func(conn net.Conn, n int) {
				rd := bufio.NewReader(conn)
				for {
					m, err := rd.ReadString('\n')
					if err != nil {
						break
					}
					msgs <- msg{fmt.Sprintf("Client %v: %v", n-1, m), conn}
				}
				// Done reading from connection
				dconns <- conn
			}(conn, n)
		case msg := <-msgs:
			// Broadcast message to conections
			for conn := range aconns {
				if msg.conn == conn {
					continue
				}
				conn.Write([]byte(msg.msg))
			}
		case dconn := <-dconns:
			log.Printf("Client %v has disconnected.\n", aconns[dconn])
			delete(aconns, dconn)
		}
	}
}
