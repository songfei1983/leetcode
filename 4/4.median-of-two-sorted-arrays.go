package medianoftwosortedarrays

import "sort"

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > 0 && len(nums2) == 0 {
		return mid(nums1)
	}
	if len(nums1) == 0 && len(nums2) > 0 {
		return mid(nums2)
	}
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	return mid(nums)
}

func mid(nums []int) float64 {
	l := len(nums)
	if l%2 == 1 {
		return float64(nums[l/2])
	}
	return float64(nums[l/2]+nums[(l-1)/2]) / 2
}

// @lc code=end
