/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */
 func subsetsWithDup(nums []int) [][]int {
	// 统计每个数出现的次数
	numCnt := make(map[int]int)
	for _, num := range nums {
		numCnt[num]++
	}
	// 记录不同的数
	numKey := make([]int, 0, len(numCnt))
	for num := range numCnt {
		numKey = append(numKey, num)
	}
	ans := make([][]int, 0, 32)
	buf := make([]int, 0, len(nums))
	dfs(&ans, &buf, numCnt, numKey, 0, )
	return ans
}

func dfs(ans *[][]int, buf *[]int, numCnt map[int]int, numKey []int, k int) {
	// 遍历完所有不同的数后退出
	if k == len(numKey) {
		newBuf := make([]int, len(*buf))
		copy(newBuf, *buf)
		*ans = append(*ans, newBuf)
		return
	}
	// 对于一个出现n次的数，有n+1次情况，分布对应添加它0-n次
	for cnt := 0; cnt <= numCnt[numKey[k]]; cnt++ {
		*buf = append(*buf, getNNumbers(numKey[k], cnt)...) // 末尾添加cnt个numkey[k]
		dfs(ans, buf, numCnt, numKey, k+1)
		*buf = (*buf)[:len(*buf)-cnt] // 移除末尾cnt个数
	}
}

func getNNumbers(num, n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = num
	}
	return nums
}
