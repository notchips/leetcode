/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */
// 先计算初始面积，然后找到最高的点之一，以此为分界点，算出两边下雨后的面积。
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	area := 0
	maxHeightIndex := 0
	for i := 0; i < n; i++ {
		area += height[i]
		if height[i] > height[maxHeightIndex] {
			maxHeightIndex = i
		}
	}

	rainArea := height[maxHeightIndex]
	for left, right := 0, 0; right < maxHeightIndex; right++ {
		if height[right] > height[left] {
			rainArea += height[left] * (right - left)
			left = right
		}
		if right+1 == maxHeightIndex {
			rainArea += height[left] * (right - left + 1)
		}
	}

	for left, right := n-1, n-1; left > maxHeightIndex; left-- {
		if height[left] > height[right] {
			rainArea += height[right] * (right - left)
			right = left
		}
		if left-1 == maxHeightIndex {
			rainArea += height[right] * (right - left + 1)
		}
	}
	
	return rainArea - area
}
