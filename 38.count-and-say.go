package main

import "strconv"

/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		s = count([]byte(s))
	}
	return s
}

func count(s []byte) string {
	n := 0
	c := s[0]
	ss := ""
	for i := 0; i < len(s); i++ {
		if c != s[i] {
			ss += strconv.Itoa(n) + string(c)
			n = 0
			c = s[i]
		}
		n++
	}
	if n > 0 {
		ss += strconv.Itoa(n) + string(s[len(s)-1:])
	}
	return ss
}

// @lc code=end
