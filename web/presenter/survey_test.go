package presenter

import (
	"math"
	"testing"
)

func Test_Survery_difference(t *testing.T) {
	type args struct {
		N float64
		P float64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "equal", args: args{N: math.Pow(10, 3), P: math.Pow(10, 3)}, want: 0},
		{name: "negative N", args: args{N: math.Pow(10, -1), P: math.Pow(10, 1)}, want: 1},
		{name: "higher N", args: args{N: math.Pow(10, 6), P: math.Pow(10, 1)}, want: 5},
		{name: "higher P", args: args{N: math.Pow(10, 1), P: math.Pow(10, 9)}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Survey{}
			if tt.args.N != 0 {
				s.N = &tt.args.N
			}
			if tt.args.P != 0 {
				s.PresurveyAssessment = &tt.args.P
			}

			if got := s.difference(); got != tt.want {
				t.Errorf("difference() = %v, want %v", got, tt.want)
			}
		})
	}
}
