package html

import "testing"

func Test_number(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "zero number", args: args{amount: 0}, want: "0"},
		{name: "positive number", args: args{amount: 15010}, want: "15,010"},
		{name: "negative number", args: args{amount: -15002}, want: "-15,002"},
		{name: "thousands number", args: args{amount: -100055}, want: "-100,055"},
		{name: "large number", args: args{amount: 123456789}, want: "123,456,789"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number(tt.args.amount); got != tt.want {
				t.Errorf("number() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_marshal(t *testing.T) {
	type args struct {
		vals any
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "empty", args: args{vals: []int{}}, want: "[]"},
		{name: "integers", args: args{vals: []int{1, 2, 3}}, want: "[1,2,3]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := marshal(tt.args.vals); string(got) != tt.want {
				t.Errorf("marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
