package convertsortedarraytobinarysearchtree

// TreeNode is struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{nums[len(nums)/2], sortedArrayToBST(nums[:len(nums)/2]), sortedArrayToBST(nums[len(nums)/2+1:])}
}

// @lc code=end
