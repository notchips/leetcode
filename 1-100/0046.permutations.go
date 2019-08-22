/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
func permute(nums []int) [][]int {
	ans := make([][]int, 0, 32)
	buf := make([]int, 0, len(nums))
	vis := make([]bool, len(nums))
	dfs(&ans, &buf, vis, nums)
	return ans
}

func dfs(ans *[][]int, buf *[]int, vis []bool, nums []int) {
	if len(*buf) == len(nums) {
		newBuf := make([]int, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !vis[i] {
			*buf = append(*buf, nums[i])
			vis[i] = true
			dfs(ans, buf, vis, nums)
			vis[i] = false
			*buf = (*buf)[:len(*buf)-1]
		}
	}
}
