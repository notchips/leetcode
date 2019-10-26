/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
package leetcode

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			l, r := i+1, len(nums)-1
			for l < r {
				if nums[l]+nums[r] == -nums[i] {
					ans = append(ans, []int{nums[i], nums[l], nums[r]})
					for l+1 < r && nums[l+1] == nums[l] {
						l++
					}
					for l < r-1 && nums[r-1] == nums[r] {
						r--
					}
					l, r = l+1, r-1
				} else if nums[l]+nums[r] < -nums[i] {
					l++
				} else {
					r--
				}
			}
		}
	}
	return ans
}

// @lc code=end
