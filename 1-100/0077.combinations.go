/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */
package leetcode

// @lc code=start
func combine(n int, k int) [][]int {
	ans := make([][]int, 0, 32)
	buf := make([]int, 0, k)
	dfs(&ans, &buf, 1, n, k)
	return ans
}

func dfs(ans *[][]int, buf *[]int, m, n, k int) {
	if len(*buf) == k {
		newBuf := make([]int, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	for i := m; i <= n; i++ {
		*buf = append(*buf, i)
		dfs(ans, buf, i+1, n, k)
		*buf = (*buf)[:len(*buf)-1]
	}
}

// @lc code=end
