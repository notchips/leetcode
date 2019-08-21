/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
 func generateParenthesis(n int) []string {
	ans := make([]string, 0, 20)
	buf := make([]byte, 0, 2*n)
	dfs(&ans, &buf, 0, 0, n)
	return ans
}

// cnt1: count of '('
// cnt2: count of ')'
func dfs(ans *[]string, buf *[]byte, cnt1, cnt2, n int) {
	if len(*buf) == 2*n {
		*ans = append(*ans, string(*buf))
		return
	}
	if cnt1 < n {
		*buf = append(*buf, '(')
		dfs(ans, buf, cnt1+1, cnt2, n)
		*buf = (*buf)[:len(*buf)-1]
	}
	if cnt2 < cnt1 {
		*buf = append(*buf, ')')
		dfs(ans, buf, cnt1, cnt2+1, n)
		*buf = (*buf)[:len(*buf)-1]
	}
}
