/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */
package leetcode

import "math"

// 用二分法快速定位两个数组的切点，切点包括数组元素的左右间隔以及元素本身
// 切点： 0  1  2  3  4
//       |  |  |  |  |
// 数组：   [1 ,  3]
// 例如切点1将数组切分为[1] 和 [1, 3]
// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) { // 选择较短数组进行二分切分
		nums1, nums2 = nums2, nums1
	}
	var leftVal1, leftVal2, rightVal1, rightVal2 int
	m, n := len(nums1), len(nums2)
	left1, right1 := 0, m*2
	for left1 <= right1 {
		mid1 := (left1 + right1) / 2
		mid2 := m + n - mid1
		leftVal1, rightVal1 = getLeftAndRightValueAtPos(nums1, mid1)
		leftVal2, rightVal2 = getLeftAndRightValueAtPos(nums2, mid2)
		if leftVal2 > rightVal1 {
			left1 = mid1 + 1
		} else if leftVal1 > rightVal2 {
			right1 = mid1 - 1
		} else {
			break
		}
	}
	return float64(max(leftVal1, leftVal2)+min(rightVal1, rightVal2)) / 2
}

// 返回切点的左右值
func getLeftAndRightValueAtPos(nums []int, pos int) (int, int) {
	n := len(nums)
	if n == 0 {
		return math.MinInt64, math.MaxInt64
	}
	switch pos {
	case 0:
		return math.MinInt64, nums[0]
	case 2 * n:
		return nums[n-1], math.MaxInt64
	default:
		return nums[(pos-1)/2], nums[pos/2]
	}
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

// @lc code=end
