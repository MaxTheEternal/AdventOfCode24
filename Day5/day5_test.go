package main

import "testing"

func TestPartOne(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			"Part One Base Example",
			args{"./day_5_input_test.txt"},
			143,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartOne(tt.args.file); got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadIncorrectPageLines(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			"Part Two Test for unorderd Page lines",
			args{"./day_5_input_test.txt"},
			135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b, c := readInput(tt.args.file)
			got := sumOfMiddlePages(listOfUpdates(a, b, c, false))
			if got != tt.want {
				t.Errorf("Wrong Lines sum %v, want %v", got, tt.want)
			}
		})
	}
}
