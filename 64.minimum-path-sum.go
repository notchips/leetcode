/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */
func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	for i := n-1; i >= 0; i--{
		for j := m-1; j >= 0; j--{
			if i < n-1 && j < m-1{
				if grid[i+1][j] < grid[i][j+1]{
					grid[i][j] += grid[i+1][j]
				} else{
					grid[i][j] += grid[i][j+1]
				}
			} else if i < n-1{
				grid[i][j] += grid[i+1][j]
			} else if j < m-1{
				grid[i][j] += grid[i][j+1]
			}
		}
	}
	return grid[0][0]
}
