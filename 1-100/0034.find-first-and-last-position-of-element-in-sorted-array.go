/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */
func searchRange(nums []int, target int) []int {
	ans := []int{-1, -1}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			ans[0] = firstEqual(nums, target, left, mid)
			ans[1] = lastEqual(nums, target, mid, right)
			break
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func firstEqual(nums []int, target, left, right int) int {
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func lastEqual(nums []int, target, left, right int) int {
	for left < right {
		mid := left + (right-left)/2 + 1
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
