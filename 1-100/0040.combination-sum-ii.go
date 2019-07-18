/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
func combinationSum2(candidates []int, target int) [][]int {
	numsCnt := make(map[int]int)
	for _, candidate := range candidates {
		numsCnt[candidate]++
	}
	nums := make([]int, 0, len(candidates))
	for num := range numsCnt {
		nums = append(nums, num)
	}

	answers := make([][]int, 0, 5)
	answer := make([]int, 0, 10)
	dfs(numsCnt, nums, &answers, &answer, 0, 0, target)
	return answers
}

func dfs(numsCnt map[int]int, nums []int, answers *[][]int, answer *[]int, i, sum, target int) {
	if sum == target {
		newRet := make([]int, len(*answer))
		copy(newRet, *answer)
		*answers = append(*answers, newRet)
		return
	}
	if sum > target || i >= len(nums) {
		return
	}

	if numsCnt[nums[i]] > 0 {
		numsCnt[nums[i]]--
		*answer = append(*answer, nums[i])
		dfs(numsCnt, nums, answers, answer, i, sum+nums[i], target)
		*answer = (*answer)[:len(*answer)-1]
		numsCnt[nums[i]]++
	}

	dfs(numsCnt, nums, answers, answer, i+1, sum, target)
}
