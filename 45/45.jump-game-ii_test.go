package jumpsgameii

import "testing"

func Test_jump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{[]int{0}}, 0},
		{"1", args{[]int{1}}, 0},
		{"2", args{[]int{1, 2}}, 1},
		{"3", args{[]int{1, 2, 3}}, 2},
		{"9", args{[]int{2, 3, 1, 1, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
