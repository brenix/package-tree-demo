package main

import (
	"reflect"
	"testing"
)

func Test_parseMsg(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   Package
		wantErr bool
	}{
		{
			name:    "Invalid package name",
			args:    args{msg: "INDEX|foo bar|\n"},
			want:    "",
			want1:   Package{},
			wantErr: true,
		},
		{
			name:    "Valid package name",
			args:    args{msg: "INDEX|foo|\n"},
			want:    "INDEX",
			want1:   Package{Name: "foo", Dependencies: []string{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := parseMsg(tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseMsg() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseMsg() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
