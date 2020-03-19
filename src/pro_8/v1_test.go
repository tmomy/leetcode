package pro_8

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{"case1", args{"  -0123456789", }, -123456789},
		//{"case2", args{"4193 with words", }, 4193},
		//{"case2", args{"3.14159", }, 3},
		//{"case3", args{"+1", }, 1},
		//{"case4", args{"010", }, 10},
		{"case4", args{"20000000000000000000"}, 2147483647},
		//{"case2", args{"PAYPALISHIRING", }, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
