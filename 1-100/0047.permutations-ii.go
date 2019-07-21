/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */
func permuteUnique(nums []int) [][]int {
	n := len(nums)
	numCnt := make(map[int]int, n)
	for i := 0; i < n; i++ {
		numCnt[nums[i]]++
	}
	answers := make([][]int, 0, 20)
	answer := make([]int, 0, n)
	dfs(&answers, &answer, numCnt, n)
	return answers
}

func dfs(answers *[][]int, answer *[]int, numCnt map[int]int, n int) {
	if len(*answer) == n {
		newAnswer := make([]int, n)
		copy(newAnswer, *answer)
		*answers = append(*answers, newAnswer)
		return
	}
	for num, cnt := range numCnt {
		if cnt > 0 {
			numCnt[num]--
			*answer = append(*answer, num)
			dfs(answers, answer, numCnt, n)
			*answer = (*answer)[:len(*answer)-1]
			numCnt[num]++
		}
	}
}
