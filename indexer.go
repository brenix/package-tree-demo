package main

import (
	"net"
)

// TODO: add description
func indexPkg(c net.Conn, r Request) {
	c.Write([]byte("INDEX FUNC\n"))
}

// TODO: add description
func removePkg(c net.Conn, r Request) {
	c.Write([]byte("REMOVE FUNC\n"))
}

// TODO: add description
func queryPkg(c net.Conn, r Request) {
	c.Write([]byte("QUERY FUNC\n"))
}
