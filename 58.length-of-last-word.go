package main

import "strings"

/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	b := []byte(strings.TrimSpace(s))
	l := len(b) - 1
	for i := l; i >= 0; i-- {
		if b[i] == ' ' {
			return l - i
		}
	}
	return len(b)
}

// @lc code=end
