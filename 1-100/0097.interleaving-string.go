/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 */
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	space := " "
	s1, s2, s3 = space+s1, space+s2, space+s3
	m, n := len(s1), len(s2)
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s2[j] == s3[i+j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s1[i] == s3[i+j]
			} else {
				dp[i][j] = (dp[i][j-1] && s2[j] == s3[i+j]) ||
					(dp[i-1][j] && s1[i] == s3[i+j])
			}
		}
	}
	return dp[m-1][n-1]
}
