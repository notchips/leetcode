/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */
package leetcode

import "sort"

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	n := len(nums)
	for i := 0; i < n-3; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			for j := i + 1; j < n-2; j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					l, r := j+1, n-1
					for l < r {
						sum := nums[i] + nums[j] + nums[l] + nums[r]
						if sum == target {
							ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
							for l+1 < r && nums[l+1] == nums[l] {
								l++
							}
							for l < r-1 && nums[r-1] == nums[r] {
								r--
							}
							l, r = l+1, r-1
						} else if sum < target {
							l++
						} else {
							r--
						}
					}
				}
			}
		}
	}
	return ans
}

// @lc code=end
