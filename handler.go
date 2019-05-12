package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Request defines the data expected from the client
type Request struct {
	Command      string
	Package      string
	Dependencies []string
}

// Response defines the response sent to the client
// type Response struct {
// 	Status string
// }

// handleConnection handles incoming connections
func handleConnection(c net.Conn) {
	log.Printf("Serving %s\n", c.RemoteAddr().String())

	for {
		d, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		r := parseRequest(d)

		// DEBUG
		if len(r.Dependencies) > 0 {
			for _, dep := range r.Dependencies {
				fmt.Println(dep)
			}
		}

		// FIXME: may want to move this into a separate function and prevent
		// redundancy
		switch r.Command {
		case "INDEX":
			go indexPkg(c, r)
		case "REMOVE":
			go removePkg(c, r)
		case "QUERY":
			go queryPkg(c, r)
		default:
			c.Write([]byte("UNKNOWN\n"))
		}

		// TODO: after further testing, need to validate this isn't disconnecting
		// clients before the command has finished. Upon initial tests, clients see
		// a EOF or disconnect
		c.Close()
		break
	}
}

// FIXME: parseRequest does not handle null / unexpected use cases
// Will need better logic to split the data and prevent unwanted data
// from corrupting the index or causing issues
//
// parseRequest takes incoming data and splits it using delimiters,
// then returns a struct of the data
func parseRequest(s string) Request {
	d := strings.Split(strings.TrimSpace(string(s)), "|")
	cmd := d[0]
	pkg := d[1]
	deps := []string{}

	for _, dep := range strings.Split(d[2], ",") {
		deps = append(deps, dep)
	}

	request := Request{Command: cmd, Package: pkg, Dependencies: deps}
	return request
}
