package pkg

import (
	"io"
	"reflect"
	"strings"
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
		results []string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name:    "should_return_error_when_results_is_empty",
			args:    args{results: []string{}},
			want:    0,
			wantErr: true,
		},
		{
			name:    "should_return_error_when_results_not_contains_expected_split_keyword",
			args:    args{results: []string{"foo", "bar"}},
			want:    0,
			wantErr: true,
		},
		{
			name:    "should_return_error_when_results_not_contains_expected_split_keyword",
			args:    args{results: []string{"foo", "1000000 bytes transferred in foo bar"}},
			want:    1000000,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTotalBytes(tt.args.results)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTotalBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getTotalBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAndPrintStdErrResult(t *testing.T) {
	type args struct {
		stderr io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "should_get_values_from_io.Read_Closer",
			args: struct {
				stderr io.ReadCloser
			}{
				stderr: io.NopCloser(strings.NewReader("foo")),
			},
			want: []string{"foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAndPrintStdErrResult(tt.args.stderr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAndPrintStdErrResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
