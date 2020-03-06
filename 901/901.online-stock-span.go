package OnlineStockSpan

/*
 * @lc app=leetcode id=901 lang=golang
 *
 * [901] Online Stock Span
 */

// @lc code=start

// StockSpanner is struct
type StockSpanner struct {
	Prices []int
	Index  int
}

// Constructor is func
func Constructor() StockSpanner {
	return StockSpanner{make([]int, 10000), 9999}
}

// Next is func
func (t *StockSpanner) Next(price int) int {
	count := 0
	t.Prices[t.Index] = price
	for i := t.Index; i < 10000; i++ {
		if price >= t.Prices[i] {
			count++
		} else {
			break
		}
	}
	t.Index--
	return count
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end
