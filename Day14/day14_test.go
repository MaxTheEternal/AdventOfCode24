package day14

import "testing"

func TestPartOne(t *testing.T) {
	type args struct {
		file    string
		seconds int
		xLen    int
		yLen    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Base Example",
			args{
				"./day14_input_test.txt",
				100,
				11,
				7,
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartOne(tt.args.file, tt.args.seconds, tt.args.xLen, tt.args.yLen); got != tt.want {
				t.Errorf("PartOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
