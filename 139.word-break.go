package main

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	memo := make([]bool, len(s))
	mwd := map[string]bool{}
	for _, v := range wordDict {
		mwd[v] = true
	}
	return check([]byte(s), mwd, 0, memo)
}

func check(s []byte, wd map[string]bool, start int, memo []bool) bool {
	if start >= len(s) {
		return true
	}
	if memo[start] {
		return true
	}
	for i := start + 1; i <= len(s); i++ {
		if _, ok := wd[string(s[start:i])]; ok && check(s, wd, i, memo) {
			memo[start] = true
			return memo[start]
		}
	}
	memo[start] = false
	return memo[start]
}

// @lc code=end
