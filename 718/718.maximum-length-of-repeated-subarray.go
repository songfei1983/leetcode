package maximumlenghtofrepeatedsubarray

/*
 * @lc app=leetcode id=718 lang=golang
 *
 * [718] Maximum Length of Repeated Subarray
 */

// @lc code=start
func findLength(A []int, B []int) int {
	ans := 0
	dp := make([][]int, len(A)+1)
	dp[0] = make([]int, len(B)+1)
	for i := 1; i <= len(A); i++ {
		dp[i] = make([]int, len(B)+1)
		for j := 1; j <= len(B); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}
	return ans
}

// func findLength(A []int, B []int) int {
// 	ans := float64(0)
// 	for i := 0; i < len(A); i++ {
// 		for j := 0; j < len(B); j++ {
// 			t := float64(0)
// 			for k := 0; i+k < len(A) && j+k < len(B) && A[i+k] == B[j+k]; k++ {
// 				t = float64(k) + 1
// 			}
// 			ans = math.Max(ans, t)
// 		}
// 	}
// 	return int(ans)
// }

// @lc code=end
