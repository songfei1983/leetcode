package carpooling

/*
 * @lc app=leetcode id=1094 lang=golang
 *
 * [1094] Car Pooling
 */

// @lc code=start
func carPooling(trips [][]int, capacity int) bool {
	var seat, walk, end int
	for {
		walk++
		if seat > capacity {
			return false
		}
		for _, v := range trips {
			if end < v[2] {
				end = v[2]
			}
			if walk == v[1] {
				seat += v[0]
			}
			if walk == v[2] {
				seat -= v[0]
			}
		}
		if walk == end {
			break
		}
	}
	return true
}

// @lc code=end
