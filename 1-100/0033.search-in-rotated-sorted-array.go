/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */
func search(nums []int, target int) int {
	n := len(nums)
	for left, right := 0, n-1; left <= right; {
		mid := (left + right) / 2
		if nums[left] <= nums[mid] { // 左半有序
			if target >= nums[left] && target <= nums[mid] {
				return binarySearch(nums, target, left, mid)
			}
			left = mid + 1
		} else { // 右半有序
			if target >= nums[mid] && target <= nums[right] {
				return binarySearch(nums, target, mid, right)
			}
			right = mid - 1
		}
	}
	return -1
}

func binarySearch(nums []int, target, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
