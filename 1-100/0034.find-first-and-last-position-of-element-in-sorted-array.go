/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */
func searchRange(nums []int, target int) []int {
	left, right := math.MaxInt32, math.MinInt32
	searchHelper(nums, target, 0, len(nums)-1, &left, &right)
	if right > -1 {
		return []int{left, right}
	}
	return []int{-1, -1}
}

func searchHelper(nums []int, target, lo, hi int, left, right *int) {
	if lo > hi {
		return
	}
	md := (lo + hi) / 2
	if target == nums[md] {
		*left = min(*left, md)
		*right = max(*right, md)
		if nums[lo] == target {
			*left = min(*left, lo)
		} else {
			searchHelper(nums, target, lo, md-1, left, right)
		}
		if nums[hi] == target {
			*right = max(*right, hi)
		} else {
			searchHelper(nums, target, md+1, hi, left, right)
		}
	} else if target < nums[md] {
		searchHelper(nums, target, lo, md-1, left, right)
	} else {
		searchHelper(nums, target, md+1, hi, left, right)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
