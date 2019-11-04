/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
package leetcode

// @lc code=start
func subsets(nums []int) [][]int {
	ans := make([][]int, 0, 32)
	buf := make([]int, 0, len(nums))
	dfs(&ans, &buf, nums)
	return ans
}

func dfs(ans *[][]int, buf *[]int, nums []int) {
	if len(nums) == 0 {
		newBuf := make([]int, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	*buf = append(*buf, nums[0])
	dfs(ans, buf, nums[1:])
	*buf = (*buf)[:len(*buf)-1]
	dfs(ans, buf, nums[1:])
}

// @lc code=end
