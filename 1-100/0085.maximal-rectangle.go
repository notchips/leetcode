/*
 * @lc app=leetcode id=85 lang=golang
 *
 * [85] Maximal Rectangle
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

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	maxArea := 0
	height := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		curRowMaxArea := largestRectangleArea(height)
		if curRowMaxArea > maxArea {
			maxArea = curRowMaxArea
		}
	}
	return maxArea
}
