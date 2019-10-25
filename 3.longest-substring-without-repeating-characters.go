package main

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	bs := []byte(s)
	max := 0
	MaxInt := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i, j := 0, 0; i < len(bs); i++ {
		if v, ok := m[bs[i]]; ok {
			j = MaxInt(j, v+1)
		}
		m[bs[i]] = i
		max = MaxInt(max, i-j+1)
	}
	return max
}

// @lc code=end
