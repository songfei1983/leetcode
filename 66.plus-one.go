package main

/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	lenArr := len(digits)
	for i := lenArr - 1; i >= 0; i-- {
		if digits[i]+1 <= 9 {
			digits[i] = digits[i] + 1
			break
		} else {
			digits[i] = 0
		}
		if i == 0 && digits[i] == 0 {
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}

// @lc code=end
