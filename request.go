package main

import (
	"log"
)

const (
	// OK response
	OK = "OK"
	// FAIL response
	FAIL = "FAIL"
	// ERROR response
	ERROR = "ERROR"
)

// handleRequest parses incoming requests, concurrently sends it to the index functions
// based on the command given, then returns a response status code/string
func handleRequest(msg string) string {
	cmd, pkg, err := parseMsg(msg)
	if err != nil {
		log.Printf("Error parsing message: %v", err)
		return ERROR
	}

	switch cmd {
	case "INDEX":
		// log.Printf("INDEX: %v", pkg.Name)
		if len(pkg.Dependencies) == 0 || inIndex(pkg.Name) {
			go addToIndex(pkg.Name, pkg.Dependencies)
			return OK // return OK when package can be indexed without dependencies or already exists
		} else if depsExist(pkg.Dependencies) {
			go addToIndex(pkg.Name, pkg.Dependencies)
			return OK // return OK when package and all of it's dependencies are indexed
		} else {
			return FAIL // return FAIL when package dependencies do not exist or package
		}
	case "QUERY":
		// log.Printf("QUERY: %v", pkg.Name)
		if inIndex(pkg.Name) {
			return OK // return OK when package is indexed
		}
		return FAIL // return FAIL when package is not indexed
	case "REMOVE":
		// log.Printf("REMOVE: %v", pkg.Name)
		if !isDep(pkg.Name) || !inIndex(pkg.Name) {
			go delFromIndex(pkg.Name)
			return OK // return OK when package can either be removed or does not exist
		}
		return FAIL // return FAIL when package still has dependencies
	default:
		return ERROR // return ERROR when command does not exist
	}
}
