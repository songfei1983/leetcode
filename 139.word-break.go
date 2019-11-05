package main

import "strings"

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	memo := map[string]bool{}
	dict := map[string]bool{}
	for _, w := range wordDict {
		dict[w] = true
	}
	return check(s, dict, memo)
}
func check(s string, dict, memo map[string]bool) bool {
	if contains(dict, s) {
		return true
	}
	if contains(memo, s) {
		return false
	}
	for w := range dict {
		if strings.HasPrefix(s, w) {
			if check(string([]byte(s)[len(w):]), dict, memo) {
				return true
			}
		}
	}
	memo[s] = false
	return false
}

func contains(d map[string]bool, s string) bool {
	_, ok := d[s]
	return ok
}

// @lc code=end
