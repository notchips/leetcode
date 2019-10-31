/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
package leetcode

// @lc code=start
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	ans := make([][]int, 0, 32)
	buf := make([]int, 0, len(nums))
	vis := make([]bool, len(nums))
	dfs46(&ans, &buf, vis, nums)
	return ans
}

func dfs46(ans *[][]int, buf *[]int, vis []bool, nums []int) {
	if len(*buf) == cap(*buf) {
		newBuf := make([]int, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	for i := range nums {
		if !vis[i] {
			*buf = append(*buf, nums[i])
			vis[i] = true
			dfs46(ans, buf, vis, nums)
			vis[i] = false
			*buf = (*buf)[:len(*buf)-1]
		}
	}
}

// @lc code=end
