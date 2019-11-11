package main

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	l0 := nums[0]
	for i := 0; i < len(nums); i++ {
		l1 := nums[i]
		subSum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			subSum += nums[j]
			if l1 <= subSum {
				l1 = subSum
			}
		}
		if l1 >= l0 {
			l0 = l1
		}
	}
	return l0
}

// @lc code=end
