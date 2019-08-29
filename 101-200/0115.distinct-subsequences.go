/*
 * @lc app=leetcode id=115 lang=golang
 *
 * [115] Distinct Subsequences
 */
/*
dp[i][j] 记录 numDistinct(s[:i+1], t[:j+1])的值
1. 当s[i] != t[i]时，s[i]不能选择，则dp[i][j] = dp[i-1][j]
2. 当s[i] == t[i]时，s[i]可选可不选（选择s[i]时只能对应t末尾字符），则dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
示例：S="babgbag", T = "bag"：
    0 1 2
    b a g
0 b 1 0 0
1 a 1 1 0
2 b 2 1 0
3 g 2 1 1
4 b 3 1 1
5 a 3 4 1
6 g 3 4 5
 */
func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m == 0 || n == 0 {
		return 0
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 确定边界
	if s[0] == t[0] {
		dp[0][0] = 1
	}
	for i := 1; i < m; i++ {
		if s[i] != t[0] {
			dp[i][0] = dp[i-1][0]
		} else {
			dp[i][0] = dp[i-1][0] + 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if s[i] != t[j] {
				dp[i][j] = dp[i-1][j]
			} else {    //  选择s[i] + 不选s[i]
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}
