/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
 func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	var (
		left, right = 0, len(nums1)*2 // nums1 left and right index
		mid1, mid2  int               // nums1, nums2 cut index
		lv1, rv1    int               // mid1 left and right value
		lv2, rv2    int               // mid2 left and right value
	)
	if len(nums1) == 0{
		mid2 = len(nums2)
		lv2, rv2 = leftAndRightValue(nums2, mid2)
		return float64(lv2 + rv2) / 2
	}
	for left <= right {
		mid1 = (left + right) / 2
		mid2 = len(nums1) + len(nums2) - mid1
		lv1, rv1 = leftAndRightValue(nums1, mid1)
		lv2, rv2 = leftAndRightValue(nums2, mid2)
		if lv1 > rv2 {
			right = mid1 - 1
		} else if lv2 > rv1 {
			left = mid1 + 1
		} else {
			break
		}
	}
	return float64(max(lv1, lv2)+min(rv1, rv2)) / 2
}

func leftAndRightValue(nums []int, mid int) (int, int) {
	if mid == 0 {
		return math.MinInt32, nums[0]
	} else if mid == 2*len(nums) {
		return nums[len(nums)-1], math.MaxInt32
	}
	return nums[(mid-1)/2], nums[mid/2]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

