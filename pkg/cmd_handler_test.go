package pkg

import (
	"testing"
)

func Test_calculateSpeed(t *testing.T) {
	type args struct {
		bytes          uint64
		sinceAsSeconds float64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "calculate", args: struct {
			bytes          uint64
			sinceAsSeconds float64
		}{bytes: 50, sinceAsSeconds: 5}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateSpeed(tt.args.bytes, tt.args.sinceAsSeconds); got != tt.want {
				t.Errorf("calculateSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTotalBytes(t *testing.T) {
	type args struct {
		args *ProgramArguments
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "should_calculate", args: struct{ args *ProgramArguments }{args: &ProgramArguments{BlockSize: 1024, Count: 1024}}, want: 1048576000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTotalBytes(tt.args.args); got != tt.want {
				t.Errorf("getTotalBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
