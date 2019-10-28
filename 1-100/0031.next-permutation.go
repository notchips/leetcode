/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */
package leetcode

import (
	"sort"
)

// @lc code=start
func nextPermutation(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] { // 存在下一个较大序列
			j := lastBiggerIndex(nums, i)
			nums[i], nums[j] = nums[j], nums[i]
			sort.Ints(nums[i+1:])
			return
		}
	}

	// 不存在下一个较大序列，说明nums整体为非递增
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

// nums[i+1:]为非递增序列，返回最后一个大于nums[i]的下标
func lastBiggerIndex(nums []int, i int) int {
	l, r := i+1, len(nums)-1
	for l < r {
		m := (l+r)/2 + 1
		if nums[m] <= nums[i] {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

// @lc code=end
