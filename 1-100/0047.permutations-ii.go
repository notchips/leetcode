/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */
func permuteUnique(nums []int) [][]int {
	numCnt := make(map[int]int)
	for _, num := range nums {
		numCnt[num]++
	}
	ans := make([][]int, 0, 32)
	buf := make([]int, 0, len(nums))
	dfs(&ans, &buf, numCnt, len(nums))
	return ans
}

func dfs(ans *[][]int, buf *[]int, numCnt map[int]int, n int) {
	if len(*buf) == n {
		newBuf := make([]int, n)
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	for num := range numCnt {
		if numCnt[num] > 0 {
			*buf = append(*buf, num)
			numCnt[num]--
			dfs(ans, buf, numCnt, n)
			numCnt[num]++
			*buf = (*buf)[:len(*buf)-1]
		}
	}
}
