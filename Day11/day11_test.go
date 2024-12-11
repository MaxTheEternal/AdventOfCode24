package day11

import (
	"reflect"
	"testing"
)

func TestPartOne(t *testing.T) {
	type args struct {
		file   string
		amount int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Part 1 Example",
			args{

				"./day11_input_test.txt",
				6,
			},
			[]string{"2097446912", "14168", "4048", "2", "0", "2", "4", "40", "48", "2024", "40", "48", "80", "96", "2", "8", "6", "7", "6", "0", "3", "2"},
		},
		{
			"Part 1 simple Test",
			args{
				"./day11_input_test_2.txt",
				1,
			},

			[]string{"1", "2024", "1", "0", "9", "9", "2021976"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartOne(tt.args.file, tt.args.amount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLength(t *testing.T) {
	type args struct {
		file   string
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Length test",
			args{
				"./day11_input_test.txt",
				6,
			},
			22,
		},
		{
			"Length Test 2",
			args{
				"./day11_input_test.txt",
				25,
			},
			55312,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := len(PartOne(tt.args.file, tt.args.amount)); got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartTwo(t *testing.T) {
	type args struct {
		file   string
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"short test",
			args{
				"./day11_input_test.txt",
				6,
			},
			22,
		},
		{
			"proper test",
			args{
				"./day11_input_test.txt",
				25,
			},
			55312,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartTwo(tt.args.file, tt.args.amount); got != tt.want {
				t.Errorf("PartTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
