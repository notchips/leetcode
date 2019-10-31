/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */
package leetcode

// @lc code=start
func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeftHeight, maxRightHeight := 0, 0
	area := 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] >= maxLeftHeight {
				maxLeftHeight = height[left]
			} else {
				area += maxLeftHeight - height[left]
			}
			left++
		} else {
			if height[right] >= maxRightHeight {
				maxRightHeight = height[right]
			} else {
				area += maxRightHeight - height[right]
			}
			right--
		}
	}
	return area
}

// 切分height，直到两端高度大于其余各点高度，这时可以直接计算出雨水面积
//func trap(height []int) int {
//	n := len(height)
//	if n < 3 {
//		return 0
//	}
//	pos := maxHeightIndex(height, 1, n-2) // 获取除两端外的最大高度索引
//	if height[pos] > height[0] || height[pos] > height[n-1] {
//		return trap(height[:pos+1]) + trap(height[pos:]) // 以最大高度为界，继续切分
//	}
//	totalArea := int(math.Min(float64(height[0]), float64(height[n-1]))) * (n - 2)
//	return totalArea - barArea(height[1:n-1])
//}
//
//func maxHeightIndex(height []int, left, right int) int {
//	pos := left
//	for i := left + 1; i <= right; i++ {
//		if height[i] > height[pos] {
//			pos = i
//		}
//	}
//	return pos
//}
//
//func barArea(height []int) int {
//	area := 0
//	for _, h := range height {
//		area += h
//	}
//	return area
//}

// @lc code=end
