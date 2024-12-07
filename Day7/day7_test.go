package day7

import (
	"testing"
)

func Test_partOne(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Part 1 Base example",
			args{
				"./day7_input_test.txt",
			},
			3749,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.args.file); got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_partTwo(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Part 2 Base Example",
			args{
				"./day7_input_test.txt",
			},
			11387,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partTwo(tt.args.file); got != tt.want {
				t.Errorf("partTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
