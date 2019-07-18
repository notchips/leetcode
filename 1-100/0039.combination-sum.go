/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
func combinationSum(candidates []int, target int) [][]int {
	rets := make([][]int, 0, 5)
	ret := make([]int, 0, 10)
	dfs(candidates, &rets, &ret, 0, 0, target)
	return rets
}

func dfs(candidates []int, rets *[][]int, ret *[]int, index, sum, target int) {
	if sum == target {
		newRet := make([]int, len(*ret))
		copy(newRet, *ret)
		*rets = append(*rets, newRet)
		return
	}
	if sum > target || index >= len(candidates) {
		return
	}

	*ret = append(*ret, candidates[index])
	dfs(candidates, rets, ret, index, sum+candidates[index], target)
	*ret = (*ret)[:len(*ret)-1]

	dfs(candidates, rets, ret, index+1, sum, target)
}
