package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

// listenAndServe creates a TCP listener and distributes client connections concurrently
func listenAndServe(port string) {
	l, err := net.Listen("tcp4", port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
		return
	}
	log.Printf("Listening on %v\n", l.Addr())

	for {
		c, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			break
		}
		go handleConn(c)
	}
}

// handleConn handles client connections to read and write responses
func handleConn(c net.Conn) {
	for {
		msg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading data from client: %v", err)
				return
			}
			break
		}
		status := handleRequest(msg)
		c.Write([]byte(status + "\n"))
	}
	// Explicitly close connection when all data has ceased or client disconnects
	c.Close()
}
