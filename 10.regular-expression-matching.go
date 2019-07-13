/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */
/*
    0 1 2 3 4 5 6 7 8 9 10 11
	_ m i s s i s s i p  p i
0 _ t f f f f f f f f f  f f
1 m f t f f f f f f f f  f f
2 i f f t f f f f f f f  f f
3 s f f f t f f f f f f  f f
4 * f f t t t f f f f f  f f
5 i f f f f f t f f f f  f f
6 s f f f f f f t f f f  f f
7 * f f f f f t t t f f  f f
8 p f f f f f f f f f f  f f
9 * f f f f f t t t f f  f f
10. f f f f f f f f t f  f f

*/
func isMatch(s string, p string) bool {
	sl, pl := len(s), len(p)
	s, p = " "+s, " "+p // 为了方便计算index，两个字符串都填充一个空字符

	dp := make([][]bool, pl+1)
	for i := 0; i <= pl; i++ {
		dp[i] = make([]bool, sl+1)
	}

	dp[0][0] = true

	for i := 1; i <= pl; i++ {
		if p[i] == '*' {
			dp[i][0] = dp[i-2][0]
		}
	}

	for i := 1; i <= pl; i++ {
		for j := 1; j <= sl; j++ {
			if p[i] == s[j] || p[i] == '.' { // match single
				dp[i][j] = dp[i-1][j-1]
			}
			if p[i] == '*' {
				if p[i-1] != s[j] && p[i-1] != '.' { // not match
					dp[i][j] = dp[i-2][j]
				} else { // match empty || match single || match more
					dp[i][j] = dp[i-2][j] || dp[i-1][j] || dp[i][j-1]
				}
			}
		}
	}

	return dp[pl][sl]
}

