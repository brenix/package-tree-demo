package main

import (
	"sync"
)

// Initialize in-memory map to index all packages and their dependencies
var index = make(map[string][]string)

// Initialize read/write mutext for locking operations
var mux = sync.RWMutex{}

// Package defines a package
type Package struct {
	Name         string
	Dependencies []string
}

// addToIndex adds a package to the index
func addToIndex(pkg string, deps []string) {
	mux.Lock()
	index[pkg] = deps
	mux.Unlock()
}

// delFromIndex deletes a package from the index
func delFromIndex(pkg string) {
	mux.Lock()
	delete(index, pkg)
	mux.Unlock()
}

// inIndex checks if a package exists in the index
func inIndex(pkg string) bool {
	mux.RLock()
	defer mux.RUnlock()
	if _, exists := index[pkg]; exists {
		return true
	}
	return false
}

// isDep checks if a package is dependent on another
func isDep(pkg string) bool {
	mux.RLock()
	defer mux.RUnlock()
	// Loop through each package in the index
	for _, deps := range index {
		// Loop through each dependency
		for dep := range deps {
			if deps[dep] == pkg {
				return true
			}
		}
	}
	return false
}

// depsExist checks that all dependencies exist in the index
func depsExist(deps []string) bool {
	mux.RLock()
	defer mux.RUnlock()
	// Loop through each dependency
	for _, dep := range deps {
		// Validate existence and return false if one does not exist
		if _, exists := index[dep]; !exists {
			return false
		}
	}
	// Return true if all depdencies were found
	return true
}
