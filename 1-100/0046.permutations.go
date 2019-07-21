/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
func permute(nums []int) [][]int {
	answers := make([][]int, 0, 20)
	answer := make([]int, 0, len(nums))
	vis := make([]bool, len(nums))
	dfs(&answers, &answer, vis, nums)
	return answers
}

func dfs(answers *[][]int, answer *[]int, vis []bool, nums []int) {
	if len(*answer) == len(nums) {
		newAnswer := make([]int, len(nums))
		copy(newAnswer, *answer)
		*answers = append(*answers, newAnswer)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !vis[i] {
			vis[i] = true
			*answer = append(*answer, nums[i])
			dfs(answers, answer, vis, nums)
			*answer = (*answer)[:len(*answer)-1]
			vis[i] = false
		}
	}
}
