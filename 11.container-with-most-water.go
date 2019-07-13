/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
 func maxArea(height []int) int {
	maxArea := 0
	for i, j := 0, len(height)-1; i < j; {
		area := min(height[i], height[j]) * (j - i)
		maxArea = max(maxArea, area)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
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

