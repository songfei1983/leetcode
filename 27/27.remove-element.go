package RemoveElement

import "fmt"

/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	n := 0
	for i, j := 0, 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[j] = nums[i]
			j++
			n++
		}
	}
	fmt.Println(nums, val, n)
	return n
}

// @lc code=end
