package day1

import (
	"testing"
)

func TestCalculateTotalDistance(t *testing.T) {
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
			args{"./day1_input_test.txt"},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateTotalDistance(tt.args.file); got != tt.want {
				t.Errorf("CalculateTotalDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimilarityScore(t *testing.T) {
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
			args{"./day1_input_test.txt"},
			31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimilarityScore(tt.args.file); got != tt.want {
				t.Errorf("SimilarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
