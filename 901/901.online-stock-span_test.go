package OnlineStockSpan

import "testing"

func TestStockSpanner_Next(t *testing.T) {
	tests := []struct {
		name  string
		price int
		want  int
	}{
		{"1", 100, 1},
		{"2", 80, 1},
		{"3", 60, 1},
		{"4", 70, 2},
		{"5", 60, 1},
		{"6", 75, 4},
		{"7", 85, 6},
	}
	this := Constructor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := this.Next(tt.price); got != tt.want {
				t.Errorf("StockSpanner.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStockSpanner_Next2(t *testing.T) {
	tests := []struct {
		name  string
		price int
		want  int
	}{
		{"1", 29, 1},
		{"2", 91, 2},
		{"3", 62, 1},
		{"4", 76, 2},
		{"5", 51, 1},
	}
	this := Constructor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := this.Next(tt.price); got != tt.want {
				t.Errorf("StockSpanner.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
