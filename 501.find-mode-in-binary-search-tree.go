package main

/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
 */
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
func find(r *TreeNode, m map[int]int) {
	if r == nil {
		return
	}
	m[r.Val]++
	find(r.Left, m)
	find(r.Right, m)
}
func findMode(root *TreeNode) []int {
	m := map[int]int{}
	max := func(m map[int]int) int {
		n := 0
		for _, v := range m {
			if v > n {
				n = v
			}
		}
		return n
	}
	find(root, m)
	maxVal := max(m)
	ns := []int{}
	for k, v := range m {
		if v == maxVal {
			ns = append(ns, k)
		}
	}
	return ns
}

// @lc code=end
