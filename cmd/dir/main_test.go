package main

import (
	"testing"
)

func BenchmarkDirShellBM(b *testing.B) {

}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"main"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"pwd", args{"."}, "main.go\nmain_test.go\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dir(tt.args.path); got != tt.want {
				t.Errorf("Dir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirTime(t *testing.T) {
	type args struct {
		args string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
        {"test", args{""}, "main.go\nmain_test.go\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirTime(tt.args.args); got != tt.want {
				t.Errorf("DirTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
