package day3

import (
	"testing"
)

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

func Test_correctMulsWithInstructions(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"part 2 test",
			args{"./day_3_input_test_part2.txt"},
			48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := correctMulsWithInstructions(tt.args.file); got != tt.want {
				t.Errorf("correctMulsWithInstructions() = %v, want %v", got, tt.want)
			}
		})
	}
}
