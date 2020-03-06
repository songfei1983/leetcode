package RemoveElement

import "testing"

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{[]int{}, 0}, 0},
		{"1", args{[]int{1}, 1}, 0},
		{"2", args{[]int{1, 1}, 1}, 0},
		{"3", args{[]int{1, 1, 2}, 1}, 1},
		{"4", args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 1}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
