/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
 // 用二分法快速定位两个数组的切点，切点包括数组元素的左右间隔以及元素本身
 // 切点： 0  1  2  3  4
 //       |  |  |  |  |
 // 数组：   [1 ,  3]
 // 切点1将数组切分为[1] 和 [1, 3]
 func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)
	var lv1, rv1, lv2, rv2 int
	if m == 0 {
		lv2, rv2 = getLeftAndRightValue(nums2, n)
		return float64(lv2+rv2) / 2
	}
	for left1, right1 := 0, 2*m; left1 <= right1; {
		mid1 := left1 + (right1-left1)/2
		mid2 := m + n - mid1
		lv1, rv1 = getLeftAndRightValue(nums1, mid1)
		lv2, rv2 = getLeftAndRightValue(nums2, mid2)
		if rv2 < lv1 {
			right1 = mid1 - 1
		} else if rv1 < lv2 {
			left1 = mid1 + 1
		} else {
			break
		}
	}
	return (math.Max(float64(lv1), float64(lv2)) + math.Min(float64(rv1), float64(rv2))) / 2
}

// 返回切点的左右值
func getLeftAndRightValue(nums []int, cutPoint int) (int, int) {
	n := len(nums)
	switch cutPoint {
	case 0:
		return math.MinInt64, nums[0]
	case 2 * n:
		return nums[n-1], math.MaxInt64
	default:
		return nums[(cutPoint-1)/2], nums[cutPoint/2]
	}
}
