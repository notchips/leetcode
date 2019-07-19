/*
 * @lc app=leetcode id=44 lang=golang
 *
 * [44] Wildcard Matching
 */
 /*
    0 1 2 3 4 5
	_ a c d c b
0 _ t f f f f f
1 a f t f f f f
2 * f t t t t t
3 c f f t f t f
4 ? f f f t f t
5 b f f f f f f
*/
func isMatch(s string, p string) bool {
	s, p = " "+s, " "+p
	sl, pl := len(s), len(p)

	dp := make([][]bool, pl)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sl)
	}

	dp[0][0] = true
	for i := 1; i < pl; i++ {
		if p[i] == '*' {
			dp[i][0] = true
		} else {
			break
		}
	}

	for i := 1; i < pl; i++ {
		for j := 1; j < sl; j++ {
			if p[i] == s[j] || p[i] == '?' { // 匹配单个字符
				dp[i][j] = dp[i-1][j-1]
			} else if p[i] == '*' { // 匹配0个 || 匹配1个 || 匹配多个
				dp[i][j] = dp[i-1][j] || dp[i-1][j-1] || dp[i][j-1]
			}
		}
	}

	return dp[pl-1][sl-1]
}
