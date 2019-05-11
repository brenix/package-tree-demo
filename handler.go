package main

import (
	"bufio"
	"errors"
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

		r, err := parseRequest(d)
		if err != nil {
			c.Write([]byte("UNKNOWN\n"))
			return
		}

		// DEBUG
		if len(r.Dependencies) > 0 {
			for _, dep := range r.Dependencies {
				fmt.Println(dep)
			}
		}

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

		c.Close()
		break
	}
}

// parseRequest takes incoming data and splits it using delimiters,
// then returns a struct of the data
func parseRequest(s string) (Request, error) {
	request := Request{}

	d := strings.Split(strings.TrimSpace(string(s)), "|")

	if len(d) < 10 {
		return request, errors.New("Null request received")
	}

	cmd := d[0]
	pkg := d[1]
	deps := []string{}

	for _, dep := range strings.Split(d[2], ",") {
		deps = append(deps, dep)
	}

	request = Request{Command: cmd, Package: pkg, Dependencies: deps}
	return request, nil
}
