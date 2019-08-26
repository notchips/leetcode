/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for row := range dp {
		dp[row] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
