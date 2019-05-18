package main

import (
	"flag"
	"fmt"
)

var port = flag.Int("port", 8080, "Port to listen for incoming connections")

// main parses command line flags and initializes a server
func main() {
	flag.Parse()
	listenAndServe(fmt.Sprintf(":%d", *port))
}
