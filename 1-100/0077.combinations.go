/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */
func combine(n int, k int) [][]int {
	ans := make([][]int, 0, 32)
	buf := make([]int, 0, k)
	dfs(&ans, &buf, 1, n, k)
	return ans
}

func dfs(ans *[][]int, buf *[]int, i, n, k int) {
	if len(*buf) == k {
		newBuf := make([]int, k)
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	if i > n {
		return
	}
	*buf = append(*buf, i)
	dfs(ans, buf, i+1, n, k)
	*buf = (*buf)[:len(*buf)-1]
	dfs(ans, buf, i+1, n, k)
}
