package sortintegersbythepowervalue

import (
	"sort"
)

/*
 * @lc app=leetcode id=1387 lang=golang
 *
 * [1387] Sort Integers by The Power Value
 */

// @lc code=start
var m map[int]int

func getKth(lo int, hi int, k int) int {
	m = map[int]int{}
	arr := make([]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		arr[i-lo] = i
	}
	sort.Slice(arr, func(i, j int) bool {
		x, y := steps(arr[i]), steps(arr[j])
		if x == y {
			return arr[i] < arr[j]
		}
		return x < y
	})
	return arr[k-1]
}

func steps(x int) int {
	if n, ok := m[x]; ok {
		return n
	}
	switch {
	case x == 1:
		m[x] = 0
	case x%2 == 1:
		m[x] = steps(x*3+1) + 1
	default:
		m[x] = steps(x/2) + 1
	}
	return m[x]
}

// @lc code=end
