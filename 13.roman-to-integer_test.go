package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"IX", args{"IX"}, 9},
		{"IV", args{"IV"}, 4},
		{"II", args{"II"}, 2},
		{"XII", args{"XII"}, 12},
		{"XXVII", args{"XXVII"}, 27},
		{"LVIII", args{"LVIII"}, 58},
		{"MCMXCIV", args{"MCMXCIV"}, 1994},
		{"MMMCMXCVIII", args{"MMMCMXCVIII"}, 3998},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
