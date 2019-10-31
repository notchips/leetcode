/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
package leetcode

// @lc code=start
func generateParenthesis(n int) []string {
	ans := make([]string, 0, 20)
	buf := make([]byte, 0, 2*n)
	dfs(&ans, &buf, n, n)
	return ans
}

func dfs(ans *[]string, buf *[]byte, l, r int) {
	if len(*buf) == cap(*buf) {
		*ans = append(*ans, string(*buf))
		return
	}
	if l > 0 {
		*buf = append(*buf, '(')
		dfs(ans, buf, l-1, r)
		*buf = (*buf)[:len(*buf)-1]
	}
	if r > 0 && r > l {
		*buf = append(*buf, ')')
		dfs(ans, buf, l, r-1)
		*buf = (*buf)[:len(*buf)-1]
	}
}

// @lc code=end
