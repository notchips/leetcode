/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
package leetcode

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var buf []int
	dfs39(candidates, target, &ans, &buf)
	return ans
}

func dfs39(candidates []int, target int, ans *[][]int, buf *[]int) {
	if target == 0 {
		newBuf := make([]int, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	if target < 0 || len(candidates) == 0 {
		return
	}
	*buf = append(*buf, candidates[0])
	dfs39(candidates, target-candidates[0], ans, buf)
	*buf = (*buf)[:len(*buf)-1]
	dfs39(candidates[1:], target, ans, buf)
}

// @lc code=end
