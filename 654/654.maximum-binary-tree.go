package maximumbinarytree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=654 lang=golang
 *
 * [654] Maximum Binary Tree
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return insert(nums, 0, len(nums)-1)
}

func maxValueIndex(nums []int, l, r int) int {
	n := l
	for i := l; i <= r; i++ {
		if nums[n] < nums[i] {
			n = i
		}
	}
	return n
}
func insert(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	maxIdx := maxValueIndex(nums, l, r)
	root := &TreeNode{Val: nums[maxIdx]}
	root.Left = insert(nums, l, maxIdx-1)
	root.Right = insert(nums, maxIdx+1, r)
	return root
}

// @lc code=end
