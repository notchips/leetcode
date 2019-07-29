/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
func subsets(nums []int) [][]int {
	var answers [][]int
	var answer []int 
	dfs(&answers, &answer, nums, 0, len(nums))
	return answers
}

func dfs(answers *[][]int, answer *[]int, nums []int, i, n int) {
	if i >= n {
		newAnswer := make([]int, len(*answer))
		copy(newAnswer, *answer)
		*answers = append(*answers, newAnswer)
		return
	}
	*answer = append(*answer, nums[i])
	dfs(answers, answer, nums, i+1, n)
	*answer = (*answer)[:len(*answer)-1]
	dfs(answers, answer, nums, i+1, n)
}
