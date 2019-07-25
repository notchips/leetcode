/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
// dp[1] = 1
// dp[2] = 2
// dp[i+2] = dp[i][i+1]
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	a, b := 1, 2
	for i := 0; i < n-2; i++ {
		a, b = b, a+b
	}
	return b
}
