package main

/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	for k, v := range nums {
		if v >= target {
			return k
		}
	}
	return len(nums)
}

// @lc code=end
