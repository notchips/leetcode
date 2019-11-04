/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 */
package leetcode

// @lc code=start
func getPermutation(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = i * factorial[i-1]
	}

	buf := make([]byte, 0, n) // 保存结果
	nums := make([]byte, n)   // 记录未选择的数字
	for i := range nums {
		nums[i] = byte(i + 1 + '0')
	}

	for k--; n > 0; n-- {
		f := factorial[n - 1]
		buf = append(buf, nums[k/f])               // 选择nums[k/f]
		nums = append(nums[:k/f], nums[k/f+1:]...) // 将nums[k/f]从nums中移除
		k %= f
	}
	return string(buf)
}

// @lc code=end
