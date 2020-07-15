package runningsumof1darray

/*
 * @lc app=leetcode id=1480 lang=golang
 *
 * [1480] Running Sum of 1d Array
 */

// @lc code=start
func runningSum(nums []int) []int {
	sums := make([]int, len(nums))
	for k, v := range nums {
		if k == 0 {
			sums[k] = v
			continue
		}
		sums[k] = sums[k-1] + v
	}
	return sums
}

// @lc code=end
