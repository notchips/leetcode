/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
package leetcode

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	numMap := make(map[int]int)
	for _, num := range candidates {
		numMap[num]++
	}
	numSet := make([]int, 0, len(numMap))
	for num := range numMap {
		numSet = append(numSet, num)
	}
	var ans [][]int
	var buf []int
	dfs(numMap, numSet, target, &ans, &buf)
	return ans
}

func dfs(numMap map[int]int, numSet []int, target int, ans *[][]int, buf *[]int) {
	if target == 0 {
		newBuf := make([]int, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	if target < 0 || len(numSet) == 0 {
		return
	}
	if num := numSet[0]; numMap[num] > 0 {
		*buf = append(*buf, num)
		numMap[num]--
		dfs(numMap, numSet, target-num, ans, buf)
		numMap[num]++
		*buf = (*buf)[:len(*buf)-1]
	}
	dfs(numMap, numSet[1:], target, ans, buf)
}

// @lc code=end
