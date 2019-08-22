/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0, 10)
	buf := make([]int, 0, 10)
	dfs(candidates, target, 0, 0, &ans, &buf)
	return ans
}

func dfs(candidates []int, target, sum, i int, ans *[][]int, buf *[]int) {
	if target == sum && len(*buf) != 0 {
		newBuf := make([]int, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	if sum > target || i >= len(candidates) {
		return
	}
	*buf = append(*buf, candidates[i])
	dfs(candidates, target, sum+candidates[i], i, ans, buf)
	*buf = (*buf)[:len(*buf)-1]
	dfs(candidates, target, sum, i+1, ans, buf)
}
