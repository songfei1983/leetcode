package main

/*
 * @lc app=leetcode id=896 lang=golang
 *
 * [896] Monotonic Array
 */

// @lc code=start
func isMonotonic(A []int) bool {
	if len(A) == 0 {
		return false
	}
	asc := func(a, b int) bool { return a <= b }
	desc := func(a, b int) bool { return a >= b }
	var check func(a, b int) bool
	bCheck := false
	for i, j := 0, 1; j < len(A); {
		if !bCheck && A[i] == A[j] {
			i++
			j++
			continue
		}
		if !bCheck && A[i] > A[j] {
			check = desc
			bCheck = true
		}
		if !bCheck && A[i] < A[j] {
			check = asc
			bCheck = true
		}
		if !check(A[i], A[j]) {
			return false
		}
		i++
		j++
	}
	return true
}

// @lc code=end
