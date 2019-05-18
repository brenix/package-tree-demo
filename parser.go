package main

import (
	"errors"
	"regexp"
	"strings"
)

// parseMsg parses a client message and returns the command, package, and error interface
func parseMsg(msg string) (string, Package, error) {
	comm := ""
	name := ""
	deps := []string{}
	pkg := Package{}

	// Validate message matches regex pattern
	regex, _ := regexp.Compile(`(\S+)\|(\S+)\|(\S+)?\.*`)
	if !regex.MatchString(msg) {
		return comm, pkg, errors.New("request does not match expected pattern")
	}

	// Split string by pipe delimiter
	s := strings.Split(strings.TrimSpace(msg), "|")

	// Assign variables
	comm = s[0]
	name = s[1]
	if len(s[2]) > 0 {
		for _, dep := range strings.Split(s[2], ",") {
			deps = append(deps, dep)
		}
	}

	// Create Package{}
	pkg = Package{Name: name, Dependencies: deps}

	return comm, pkg, nil
}
