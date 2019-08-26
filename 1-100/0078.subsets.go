/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
func subsets(nums []int) [][]int {
	ans := make([][]int, 0, 32)
	buf := make([]int, 0, len(nums))
	dfs(&ans, &buf, nums, 0, len(nums))
	return ans
}

func dfs(ans *[][]int, buf *[]int, nums []int, i, n int) {
	if i == n {
		newBuf := make([]int, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	*buf = append(*buf, nums[i])
	dfs(ans, buf, nums, i+1, n)
	*buf = (*buf)[:len(*buf)-1]
	dfs(ans, buf, nums, i+1, n)
}
