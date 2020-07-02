package interleavingstring

/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	if len(s1) == 0 || len(s2) == 0 {
		return (s1 == s3) || (s2 == s3)
	}
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
		for j := range dp[i] {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && []rune(s2)[j-1] == []rune(s3)[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && []rune(s1)[i-1] == []rune(s3)[i+j-1]
			} else {
				dp[i][j] = (dp[i-1][j] && []rune(s1)[i-1] == []rune(s3)[i+j-1]) || dp[i][j-1] && []rune(s2)[j-1] == []rune(s3)[i+j-1]
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

// @lc code=end
