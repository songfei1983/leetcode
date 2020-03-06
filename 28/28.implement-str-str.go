package ImplementStrStr

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if haystack == needle || needle == "" {
		return 0
	}
	bs := []byte(haystack)
	x, y := len(bs), len(needle)
	for i := 0; i <= x-y; i++ {
		if string(bs[i:i+y]) == needle {
			return i
		}
	}
	return -1
}

// @lc code=end
