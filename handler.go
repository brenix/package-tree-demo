package main

import (
	"bufio"
	"errors"
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
		// Use func for defer to avoid issues in loop
		defer func() {
			c.Close()
		}()

		d, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		r, err := parseRequest(d)
		if err != nil {
			log.Println(err)
			c.Write([]byte("UNKNOWN\n"))
			return
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
	}
}

// FIXME: parseRequest does not handle null / unexpected use cases
// Will need better logic to split the data and prevent unwanted data
// from corrupting the index or causing issues
//
// parseRequest takes incoming data and splits it using delimiters,
// then returns a struct of the data
func parseRequest(s string) (Request, error) {
	r := Request{}
	cmd := ""
	pkg := ""
	deps := []string{}

	d := strings.Split(strings.TrimSpace(string(s)), "|")

	if len(d[0]) > 0 && len(d[1]) > 0 {
		cmd = d[0]
		pkg = d[1]
	} else {
		return r, errors.New("Unable to parse request: null value detected")
	}

	for _, dep := range strings.Split(d[2], ",") {
		deps = append(deps, dep)
	}

	r = Request{Command: cmd, Package: pkg, Dependencies: deps}
	return r, nil
}
