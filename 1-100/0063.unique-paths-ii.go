/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else if i == n-1 && j == m-1 {
				dp[i][j] = 1
			} else if i == n-1 {
				dp[i][j] = dp[i][j+1]
			} else if j == m-1 {
				dp[i][j] = dp[i+1][j]
			} else {
				dp[i][j] = dp[i][j+1] + dp[i+1][j]
			}
		}
	}

	return dp[0][0]
}

