/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */
/*
    0 1 2 3 4 5 6 7 8 9 A B
    _ m i s s i s s i p p i
0 _ t f f f f f f f f f f f
1 m f t f f f f f f f f f f
2 i f f t f f f f f f f f f
3 s f f f t f f f f f f f f
4 * f f t t t f f f f f f f
5 i f f f f f t f f f f f f
6 s f f f f f f t f f f f f
7 * f f f f f t t t f f f f
8 p f f f f f f f f f f f f
9 * f f f f f t t t f f f f
A . f f f f f f t t t f f f
*/
func isMatch(s string, p string) bool {
	s, p = " "+s, " "+p
	m, n := len(p), len(s)
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	// 初始化第一列
	dp[0][0] = true
	for i := 2; i < m; i++ {
		if p[i] == '*' {
			dp[i][0] = dp[i-2][0]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if p[i] == '.' || p[i] == s[j] { // 匹配单个字符相等
				dp[i][j] = dp[i-1][j-1]
			} else if p[i] == '*' {
				if p[i-1] != '.' && p[i-1] != s[j] { // 匹配失败
					dp[i][j] = dp[i-2][j]
				} else if i >= 2 { //匹配0个，匹配1个，匹配多个
					dp[i][j] = dp[i-2][j] || dp[i-1][j] || dp[i][j-1]
				}
			}
		}
	}
	return dp[m-1][n-1]
}
