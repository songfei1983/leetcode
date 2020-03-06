package SqrtX

/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	a := float64(x)
	b := float64(a) / 2
	i := 100
	for {
		b = (b + a/b) / 2.0
		if i < 0 {
			break
		}
		i--
	}

	return int(b)
}

// @lc code=end
