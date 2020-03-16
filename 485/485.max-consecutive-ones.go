package maxconsecutiveones

/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max, cnt := 0, 0
	for _, v := range nums {
		cnt += v
		if v == 0 {
			if cnt > max {
				max, cnt = cnt, 0
			}
			cnt = 0
		}
	}
	if cnt > max {
		max, cnt = cnt, 0
	}
	return max
}

// @lc code=end
