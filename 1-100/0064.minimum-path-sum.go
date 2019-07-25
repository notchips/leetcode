/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */
 func minPathSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i != n-1 && j != m-1 {
				grid[i][j] += min(grid[i][j+1], grid[i+1][j])
			} else if i != n-1 {
				grid[i][j] += grid[i+1][j]
			} else if j != m-1 {
				grid[i][j] += grid[i][j+1]
			}
		}
	}
	return grid[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
