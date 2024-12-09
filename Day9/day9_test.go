package day9

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
			"Base Example",
			args{
				"./day9_input_test.txt",
			},
			1928,
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
