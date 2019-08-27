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
		} else if nums[left] < nums[mid] { //右半有序
			if nums[left] <= target && target <= nums[mid] {
				return binarySearch(nums[left:mid+1], target)
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[right] { // 左半有序
			if nums[mid] <= target && target <= nums[right] {
				return binarySearch(nums[mid:right+1], target)
			} else {
				right = mid - 1
			}
		} else { // 无法判断左右是否有序
			return search(nums[left:mid], target) || search(nums[mid+1:right+1], target)
		}
	}
	return false
}

func binarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
