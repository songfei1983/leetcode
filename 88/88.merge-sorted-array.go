package mergesortedarray

import "sort"

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n > 0 {
		for i := m; i < m+n; i++ {
			nums1[i] = nums2[i-m]
		}
	}
	sort.Ints(nums1)
}

// @lc code=end
