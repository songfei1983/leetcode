package main

/*
 * @lc app=leetcode id=901 lang=golang
 *
 * [901] Online Stock Span
 */

// @lc code=start
type StockSpanner struct {
	Prices []int
	Index  int
}

func Constructor() StockSpanner {
	return StockSpanner{make([]int, 10000), 9999}
}

func (this *StockSpanner) Next(price int) int {
	count := 0
	this.Prices[this.Index] = price
	for i := this.Index; i < 10000; i++ {
		if price >= this.Prices[i] {
			count++
		} else {
			break
		}
	}
	this.Index--
	return count
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
// @lc code=end
