/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */
package leetcode

// @lc code=start
func permuteUnique(nums []int) [][]int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	numSet := make([]int, 0, len(nums))
	for num := range numMap {
		numSet = append(numSet, num)
	}

	ans := make([][]int, 0, 32)
	buf := make([]int, 0, len(nums))
	dfs(&ans, &buf, numMap, numSet)
	return ans
}

func dfs(ans *[][]int, buf *[]int, numMap map[int]int, numSet []int) {
	if len(*buf) == cap(*buf) {
		newBuf := make([]int, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	for _, num := range numSet {
		if numMap[num] > 0 {
			*buf = append(*buf, num)
			numMap[num]--
			dfs(ans, buf, numMap, numSet)
			numMap[num]++
			*buf = (*buf)[:len(*buf)-1]
		}
	}
}

// @lc code=end
