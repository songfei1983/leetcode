package main

/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	n := map[byte]int{}
	for _, v := range []byte(s) {
		m[v]++
		if m[v] > 1 {
			n[v] = 1
		}
	}
	return len(n)
}

// @lc code=end
