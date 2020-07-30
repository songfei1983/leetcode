package carpooling

import "testing"

func Test_carPooling(t *testing.T) {
	type args struct {
		trips    [][]int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{trips: [][]int{{2, 1, 5}, {3, 3, 7}}, capacity: 4}, false},
		{"2", args{trips: [][]int{{2, 1, 5}, {3, 3, 7}}, capacity: 5}, true},
		{"3", args{trips: [][]int{{2, 1, 5}, {3, 5, 7}}, capacity: 3}, true},
		{"4", args{trips: [][]int{{3, 2, 7}, {3, 7, 9}, {8, 3, 9}}, capacity: 11}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carPooling(tt.args.trips, tt.args.capacity); got != tt.want {
				t.Errorf("carPooling() = %v, want %v", got, tt.want)
			}
		})
	}
}
