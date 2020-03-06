package LongestCommonPrefix

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"6", args{[]string{"aaa", "aa", "aaa"}}, "aa"},
		{"1", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"2", args{[]string{"dog", "racecar", "car"}}, ""},
		{"3", args{[]string{"", "b"}}, ""},
		{"4", args{[]string{"a"}}, "a"},
		{"5", args{[]string{"aa", "a"}}, "a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
