package maximumdepthofbinarytree

// TreeNode is struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	return max(l, r) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
