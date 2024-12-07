package day6

import (
	"testing"
)

func Test_partOne(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Part One Bae Example",
			args{
				"./day6_input_test.txt",
			},
			41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.args.path); got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isLooping(t *testing.T) {
	type args struct {
		matrix   [][]string
		x        int
		y        int
		xBlocker int
		yBlocker int
	}
	matrix, x, y := readFile("./day6_input_test.txt")
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test looping 1",
			args{matrix, x, y, 3, 6},
			true,
		},
		{
			"test looping 2",
			args{matrix, x, y, 6, 7},
			true,
		},
		{
			"test looping 3",
			args{matrix, x, y, 7, 7},
			true,
		},
		{
			"test looping 4",
			args{matrix, x, y, 1, 8},
			true,
		},
		{
			"test looping 5",
			args{matrix, x, y, 3, 8},
			true,
		},
		{
			"test looping 6",
			args{matrix, x, y, 7, 9},
			true,
		},
		{
			"test not looping",
			args{matrix, x, y, -1, -1},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLooping(tt.args.matrix, tt.args.x, tt.args.y, tt.args.xBlocker, tt.args.yBlocker); got != tt.want {
				t.Errorf("isLooping() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestPartTwo(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Part 2 Test",
			args{
				"./day6_input_test.txt",
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.path); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
