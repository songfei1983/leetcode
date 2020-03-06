package LongestCommonPrefix

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	switch n {
	case 0:
		return ""
	case 1:
		return strs[0]
	}
	prefix := []byte(strs[0])
	for _, str := range strs[1:] {
		if len(prefix) == 0 || len(str) == 0 {
			return ""
		}
		if len(prefix) > len(str) {
			prefix = prefix[:len(str)]
		}
		for k, v := range []byte(str) {
			if k >= len(prefix) {
				break
			}
			if prefix[k] != v {
				prefix = prefix[:k]
				break
			}
		}
	}
	return string(prefix)
}

// @lc code=end
