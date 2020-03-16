package scoreafterflippingmatrix

/*
 * @lc app=leetcode id=861 lang=golang
 *
 * [861] Score After Flipping Matrix
 */

// @lc code=start
func matrixScore(A [][]int) int {
	n, m := len(A), len(A[0])
	for i := 0; i < n; i++ {
		if A[i][0] == 1 {
			continue
		}
		for j := 0; j < m; j++ {
			A[i][j] ^= 1
		}
	}
	res := (1 << (m - 1)) * n
	for j := 1; j < m; j++ {
		cnt := 0
		for i := 0; i < n; i++ {
			cnt += A[i][j]
		}
		if cnt < (n - cnt) {
			cnt = n - cnt
		}
		res += (1 << (m - 1 - j)) * cnt
	}
	return res
}

// @lc code=end
