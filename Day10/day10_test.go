package day10

import (
	"testing"
)

func TestCalcPaths(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{

			"Part 1 Base Example",
			args{
				"./day10_input_test.txt",
			},
			36,
			81,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CalcPaths(tt.args.file)
			if got != tt.want {
				t.Errorf("CalcPaths() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CalcPaths() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
