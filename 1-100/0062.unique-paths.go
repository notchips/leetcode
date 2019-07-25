/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
func uniquePaths(m int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i == n-1 || j == m-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j+1] + dp[i+1][j]
			}
		}
	}

	return dp[0][0]
}

