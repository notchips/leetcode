/*
 * @lc app=leetcode id=84 lang=golang
 *
 * [84] Largest Rectangle in Histogram
 */
func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	maxArea := 0
	leftFirstLessIndex := make([]int, n)
	leftFirstLessIndex[0] = -1
	for i := 1; i < n; i++ {
		left := i - 1
		for left >= 0 && heights[left] >= heights[i] {
			left = leftFirstLessIndex[left]
		}
		leftFirstLessIndex[i] = left
	}

	rightFirstLessIndex := make([]int, n)
	rightFirstLessIndex[n-1] = n
	for i := n - 2; i >= 0; i-- {
		right := i + 1
		for right < n && heights[right] >= heights[i] {
			right = rightFirstLessIndex[right]
		}
		rightFirstLessIndex[i] = right
	}

	for i := 0; i < n; i++ {
		area := heights[i] * (rightFirstLessIndex[i] - leftFirstLessIndex[i] - 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
