package surfaceareaof3dshapes

/*
 * @lc app=leetcode id=892 lang=golang
 *
 * [892] Surface Area of 3D Shapes
 */

// @lc code=start
func surfaceArea(grid [][]int) int {
	// sum, 垂直重叠, 行重叠, 列重叠
	sum, verticalOverlap, rowOverlap, colOverlap := 0, 0, 0, 0
	for i := range grid {
		for j := range grid[i] {
			sum += grid[i][j]
			if grid[i][j] > 1 {
				verticalOverlap += grid[i][j] - 1
			}
			if j > 0 {
				rowOverlap += min(grid[i][j-1], grid[i][j])
			}
			if i > 0 {
				colOverlap += min(grid[i-1][j], grid[i][j])
			}
		}
	}
	return sum*6 - (verticalOverlap+rowOverlap+colOverlap)*2
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
