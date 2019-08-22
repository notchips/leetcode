/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
func combinationSum2(candidates []int, target int) [][]int {
	numCnt := make(map[int]int)
	for _, num := range candidates {
		numCnt[num]++
	}
	nums := make([]int, 0, len(numCnt))
	for num := range numCnt {
		nums = append(nums, num)
	}
	ans := make([][]int, 0, 10)
	buf := make([]int, 0, 10)
	dfs(numCnt, nums, target, 0, 0, &ans, &buf)
	return ans
}

func dfs(numCnt map[int]int, nums []int, target, sum, i int, ans *[][]int, buf *[]int) {
	if target == sum && len(*buf) != 0 {
		newBuf := make([]int, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	if sum > target || i >= len(nums) {
		return
	}
	if numCnt[nums[i]] > 0 {
		*buf = append(*buf, nums[i])
		numCnt[nums[i]]--
		dfs(numCnt, nums, target, sum+nums[i], i, ans, buf)
		*buf = (*buf)[:len(*buf)-1]
		numCnt[nums[i]]++
	}
	dfs(numCnt, nums, target, sum, i+1, ans, buf)
}
