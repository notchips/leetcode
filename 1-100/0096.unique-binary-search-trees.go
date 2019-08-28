/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */
// dp[0] = 1
// dp[1] = 1
// dp[2] = dp[0]*dp[1] + dp[1]*dp[0] = 2
// dp[3] = dp[0]*dp[2] + dp[1]*dp[1] + dp[2]*dp[0] = 5
// dp[4] = dp[0]*dp[3] + dp[1]*dp[2] + dp[2]*dp[1] + dp[3]*dp[0] = 14
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j <= (i-1)/2; j++ {
			if j == i-1-j {
				dp[i] += dp[j] * dp[j]
			} else {
				dp[i] += 2 * dp[j] * dp[i-1-j]
			}
		}
	}
	return dp[n]
}
