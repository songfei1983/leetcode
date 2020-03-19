package binarytreelevelordertraversalii

// TreeNode is struce
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	level := 0
	find(root, level, &res)
	reverse(&res)
	return res
}

func reverse(res *[][]int) {
	for i, j := 0, len(*res)-1; i < j; {
		(*res)[i], (*res)[j] = (*res)[j], (*res)[i]
		i++
		j--
	}
}

func find(t *TreeNode, level int, res *[][]int) {
	if t == nil {
		return
	}
	if len(*res) > level {
		(*res)[level] = append((*res)[level], t.Val)
	} else {
		*res = append(*res, []int{t.Val})
	}
	find(t.Left, level+1, res)
	find(t.Right, level+1, res)
}

// @lc code=end
