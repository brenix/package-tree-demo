package main

import "testing"

func Test_handleRequest(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Invalid index",
			args: args{"IDNEX|foo|\n"},
			want: "ERROR",
		},
		{
			name: "Valid remove",
			args: args{"REMOVE|foo|\n"},
			want: "OK",
		},
		{
			name: "Null",
			args: args{""},
			want: "ERROR",
		},
		{
			name: "Malformed index",
			args: args{"INDEX|foo=bar|baz"},
			want: "FAIL",
		},
		{
			name: "Valid query",
			args: args{"QUERY|qux|\n"},
			want: "FAIL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleRequest(tt.args.msg); got != tt.want {
				t.Errorf("handleRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
