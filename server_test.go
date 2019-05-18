package main

import (
	"net"
	"testing"
)

func Test_listenAndServe(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Run Server",
			args: args{address: "127.0.0.1:8080"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go listenAndServe(tt.args.address)
		})
	}
}

func Test_handleConn(t *testing.T) {
	_, client := net.Pipe()
	defer client.Close()

	type args struct {
		c net.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Client connection",
			args: args{client},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go handleConn(tt.args.c)
		})
	}
}
