/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */
func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[right] > nums[mid] { // 右半有序
			if binSearch(nums, target, mid+1, right) {
				return true
			}
			right = mid - 1
		} else if nums[left] < nums[mid] { // 左半有序
			if binSearch(nums, target, left, mid-1) {
				return true
			}
			left = mid + 1
		} else { // 无法判断左右是否有序
			return search(nums[left:mid], target) || search(nums[mid+1:right+1], target)
		}
	}
	return false
}

func binSearch(nums []int, target, left, right int) bool {
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
