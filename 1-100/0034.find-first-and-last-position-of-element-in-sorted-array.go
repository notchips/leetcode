/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */
package leetcode

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{findFirstIndex(nums, target), findLastIndex(nums, target)}
}

func findFirstIndex(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func findLastIndex(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l+r)/2 + 1
		if nums[m] > target {
			r = m - 1
		} else {
			l = m
		}
	}
	if nums[l] == target {
		return l
	}
	return -1

}

// @lc code=end
