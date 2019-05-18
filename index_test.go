package main

import "testing"

func Test_addToIndex(t *testing.T) {
	type args struct {
		pkg  string
		deps []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Add package",
			args: args{pkg: "foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addToIndex(tt.args.pkg, tt.args.deps)
		})
	}
}

func Test_delFromIndex(t *testing.T) {
	type args struct {
		pkg string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Delete package",
			args: args{pkg: "foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			delFromIndex(tt.args.pkg)
		})
	}
}

func Test_inIndex(t *testing.T) {
	type args struct {
		pkg string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Package does not exist",
			args: args{pkg: "foo"},
			want: false,
		},
		{
			name: "Package exists",
			args: args{pkg: "bar"},
			want: true,
		},
	}
	addToIndex("bar", []string{})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inIndex(tt.args.pkg); got != tt.want {
				t.Errorf("inIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isDep(t *testing.T) {
	type args struct {
		pkg string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Package is not a dependency",
			args: args{pkg: "foo"},
			want: false,
		},
		{
			name: "Package is a dependency",
			args: args{pkg: "bar"},
			want: true,
		},
	}
	index["foo"] = []string{"bar"}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDep(tt.args.pkg); got != tt.want {
				t.Errorf("isDep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_depsExist(t *testing.T) {
	type args struct {
		deps []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Dependencies do not exist",
			args: args{[]string{}},
			want: true,
		},
		{
			name: "Dependencies exist",
			args: args{[]string{"baz"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := depsExist(tt.args.deps); got != tt.want {
				t.Errorf("depsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
