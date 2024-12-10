package day10

import "testing"

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

			"Part 1 Base Example",
			args{
				"./day10_input_test.txt",
			},
			36,
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
