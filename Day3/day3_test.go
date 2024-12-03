package main

import "testing"

func Test_correctMuls(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"base example",
			args{"./day3_input_test.txt"},
			161,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := correctMuls(tt.args.file); got != tt.want {
				t.Errorf("correctMuls() = %v, want %v", got, tt.want)
			}
		})
	}
}
