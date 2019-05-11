package main

import (
	"log"
	"net"
)

// listenAndServe creates a TCP listener and serves client connections
// concurrently using goroutines
func listenAndServe(port string) {
	l, err := net.Listen("tcp4", port)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Listening on %v\n", l.Addr())
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go handleConnection(c)
	}
}
