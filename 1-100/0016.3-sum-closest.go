/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */
package leetcode

import (
	"sort"
)

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	ans := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			l, r := i+1, len(nums)-1
			for l < r {
				sum := nums[i] + nums[l] + nums[r]
				if sum == target {
					return target
				} else if sum < target {
					l++
				} else {
					r--
				}
				if abs(sum-target) < abs(ans-target) {
					ans = sum
				}
			}
		}
	}
	return ans
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// @lc code=end
