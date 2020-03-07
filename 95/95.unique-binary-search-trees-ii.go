package uniquebinarysearchtreeii

import (
	"strconv"
)

/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 */

// TreeNode is struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	ts := make([]*TreeNode, 0)
	if n == 1 {
		return append(ts, &TreeNode{Val: n})
	}
	m := map[string]bool{}
	Perm(n, func(a []int) {
		node := new(TreeNode)
		node.Val = a[0]
		for i := 1; i < len(a); i++ {
			insert(node, a[i])
		}
		var key string
		preOrderTraverse(node, func(i int) {
			key += strconv.Itoa(i)
		})
		if _, ok := m[key]; !ok {
			m[key] = true
			ts = append(ts, node)
		}
	})
	return ts
}

func preOrderTraverse(n *TreeNode, f func(i int)) {
	if n != nil {
		f(n.Val)
		preOrderTraverse(n.Left, f)
		preOrderTraverse(n.Right, f)
	}
}

func insert(n *TreeNode, data int) {
	if n == nil {
		return
	} else if data <= n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: data, Left: nil, Right: nil}
		} else {
			insert(n.Left, data)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Val: data, Left: nil, Right: nil}
		} else {
			insert(n.Right, data)
		}
	}
}

// Perm calls f with each permutation of a.
func Perm(n int, f func([]int)) {
	a := make([]int, n)
	for k := range a {
		a[k] = k + 1
	}
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// @lc code=end
