/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */
package leetcode

// @lc code=start
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] <= nums[r] { // 右边有序
			if nums[m] <= target && target <= nums[r] {
				return binarySearch(nums, m, r, target)
			}
			r = m - 1
		} else { // 左边有序
			if nums[l] <= target && target <= nums[m] {
				return binarySearch(nums, l, m, target)
			}
			l = m + 1
		}
	}
	return -1
}

func binarySearch(nums []int, l, r, target int) int {
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

// @lc code=end
