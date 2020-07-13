package sortintegersbythepowervalue

import "testing"

func Test_getKth(t *testing.T) {
	type args struct {
		lo int
		hi int
		k  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{12, 15, 2}, 13},
		{"2", args{1, 1, 1}, 1},
		{"3", args{7, 11, 4}, 7},
		{"4", args{10, 20, 5}, 13},
		{"5", args{1, 1000, 777}, 570},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKth(tt.args.lo, tt.args.hi, tt.args.k); got != tt.want {
				t.Errorf("getKth() = %v, want %v", got, tt.want)
			}
		})
	}
}
