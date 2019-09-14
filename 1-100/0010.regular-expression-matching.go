/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */
package leetcode

// @lc code=start
func isMatch(s string, p string) bool {
	s, p = " "+s, " "+p
	m, n := len(p), len(s)
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	for i := 2; i < m; i++ {
		if p[i] == '*' {
			dp[i][0] = dp[i-2][0]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if p[i] == '.' || p[i] == s[j] { // 匹配单个字符
				dp[i][j] = dp[i-1][j-1]
			} else if p[i] == '*' && i > 1 {
				if p[i-1] == '.' || p[i-1] == s[j] { //匹配0个，匹配1个，匹配多个
					dp[i][j] = dp[i-2][j] || dp[i-1][j-1] || dp[i][j-1]
				} else { // 匹配0个
					dp[i][j] = dp[i-2][j]
				}
			}
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end
