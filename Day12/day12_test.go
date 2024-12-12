package day12

import (
	"testing"
)

func Test_calcFencePrices(t *testing.T) {
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
				"./day12_input_test.txt",
			},
			1930,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFencePrices(tt.args.file); got != tt.want {
				t.Errorf("calcFencePrices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcWithBulkDiscount(t *testing.T) {
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
				"./day12_input_test.txt",
			},
			1206,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcWithBulkDiscount(tt.args.file); got != tt.want {
				t.Errorf("calcWithBulkDiscount() = %v, want %v", got, tt.want)
			}
		})
	}
}
