package LongestPalindromicSubstring

import (
	"strings"
)

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	ch := make([]byte, 0)
	ch = append(ch, '#')
	for _, v := range s {
		ch = append(ch, byte(v), '#')
	}
	maxLen := 0
	chs := []byte{}
	for k := range ch {
		if x := manache(ch, k); x > maxLen {
			maxLen = x
			chs = ch[k-maxLen : k+maxLen]
		}
	}
	return strings.ReplaceAll(string(chs), "#", "")
}

func manache(ch []byte, i int) (l int) {
	for k := 0; k < i && k < len(ch)-i; k++ {
		if ch[i-k] == ch[i+k] {
			l++
		} else {
			break
		}
	}
	if i+l < len(ch) && ch[i+l] == '#' {
		return l
	}
	return l - 1
}

// @lc code=end
