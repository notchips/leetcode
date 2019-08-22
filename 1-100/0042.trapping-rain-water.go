/*
 * @lc app=leetcode id=42 lang=golang
 *
 * [42] Trapping Rain Water
 */
// 切分height，直到两端高度大于其余各点高度，这时可以直接计算出雨水面积
func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}
	pos := maxHeightIndex(height, 1, n-2) // 获取除两端外的最大高度索引
	if height[pos] <= height[0] && height[pos] <= height[n-1] { // 两端高度大于其余各点高度，直接得出雨水面积
		return min(height[0], height[n-1])*(n-2) - area(height[1:n-1])
	}
	return trap(height[0:pos+1]) + trap(height[pos:]) // 以最大高度为界，继续切分
}

func maxHeightIndex(height []int, left, right int) int {
	pos := left
	for i := left + 1; i <= right; i++ {
		if height[i] > height[pos] {
			pos = i
		}
	}
	return pos
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func area(height []int) int {
	sum := 0
	for i := 0; i < len(height); i++ {
		sum += height[i]
	}
	return sum
}
