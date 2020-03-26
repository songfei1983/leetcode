package balancedbinarytree

import "math"

// TreeNode is struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	return root == nil ||
		isBalanced(root.Left) &&
			math.Abs(height(root.Left)-height(root.Right)) < 2 &&
			isBalanced(root.Right)
}

func height(t *TreeNode) float64 {
	if t == nil {
		return 0
	}
	return math.Max(height(t.Left), height(t.Right)) + 1
}

// @lc code=end
