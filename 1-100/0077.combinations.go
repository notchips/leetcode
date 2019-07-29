/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */
func combine(n int, k int) [][]int {
	var answers [][]int
	var answer []int
	dfs(&answers, &answer, 1, n, k)
	return answers
}

func dfs(answers *[][]int, answer *[]int, start, n, k int) {
	if len(*answer) == k {
		newAnswer := make([]int, k)
		copy(newAnswer, *answer)
		*answers = append(*answers, newAnswer)
		return
	}
	for i := start; i <= n; i++ {
		*answer = append(*answer, i)
		dfs(answers, answer, i+1, n, k)
		*answer = (*answer)[:len(*answer)-1]
	}
}
