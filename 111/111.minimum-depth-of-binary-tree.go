package minimumdepthofbinarytree

import "math"

/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if (root.Left == nil) && (root.Right == nil) {
		return 1
	}
	d := math.MaxInt16
	if root.Left != nil {
		d = int(math.Min(float64(minDepth(root.Left)), float64(d)))
	}
	if root.Right != nil {
		d = int(math.Min(float64(minDepth(root.Right)), float64(d)))
	}
	return d + 1
}

// @lc code=end

// TreeNode is struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
