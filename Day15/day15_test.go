package day15

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
			"Large Example",
			args{
				"./day15_input_test.txt",
			},
			10092,
		},
		{
			"Small Example",
			args{
				"./day15_input_test_small.txt",
			},
			2028,
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
