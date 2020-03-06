package LengthOfLastWord

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"hello world"}, 5},
		{"2", args{""}, 0},
		{"3", args{"song fei"}, 3},
		{"4", args{"song fei "}, 3},
		{"5", args{"song fei  "}, 3},
		{"6", args{"    "}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
