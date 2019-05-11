package main

import (
	"flag"
	"fmt"
)

var port = flag.Int("port", 8080, "Port to listen for incoming connections")

func main() {
	flag.Parse()
	listenAndServe(fmt.Sprintf(":%d", *port))
}
