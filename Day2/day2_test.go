package day2

import (
	"testing"
)

func TestSafeReports(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"base example",
			args{
				"./day2_input_test.txt",
			},
			4,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SafeReports(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("SafeReports() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SafeReports() = %v, want %v", got, tt.want)
			}
		})
	}
}
