package ImplementStrStr

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"hello", "ll"}, 2},
		{"2", args{"aaaaa", "bba"}, -1},
		{"3", args{"", ""}, 0},
		{"4", args{"x", ""}, 0},
		{"5", args{"", "y"}, -1},
		{"6", args{"a", "a"}, 0},
		{"7", args{"mississippi", "pi"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
