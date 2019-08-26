/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */
func minDistance(word1 string, word2 string) int {
	word1, word2 = " "+word1, " "+word2
	m, n := len(word1), len(word2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = i
	}
	for j := 0; j < n; j++ {
		dp[0][j] = j
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else { // delete , replace, insert
				dp[i][j] = min(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]+1)
			}
		}
	}
	return dp[m-1][n-1]
}

func min(a, b, c int) int {
	if b < a {
		a, b = b, a
	}
	if c < a {
		a, c = c, a
	}
	return a
}
